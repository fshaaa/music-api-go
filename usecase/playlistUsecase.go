package usecase

import (
	"music-api-go/dto"
	"music-api-go/model"
	"music-api-go/repository/playlistSongsRepository"
	"music-api-go/repository/playlistsRepository"
)

type PlaylistUsecase interface {
	AddPlaylist(playlist model.Playlists) error
	DeletePlaylist(id string) error
	GetPlaylistByID(id string) (dto.Playlist, error)
	GetAllPlaylists() ([]dto.Playlist, error)
	GetPlaylistDetail(id string) (dto.PlaylistDetail, error)
}

type playlistUsecase struct {
	playlist     playlistsRepository.PlaylistsRepository
	playlistSong playlistSongsRepository.PlaylistSongsRepository
	songUC       *songUsecase
	userUC       *userUsecase
}

func NewPlaylistUsecase(p playlistsRepository.PlaylistsRepository, ps playlistSongsRepository.PlaylistSongsRepository,
	sc *songUsecase, uc *userUsecase) *playlistUsecase {
	return &playlistUsecase{p, ps, sc, uc}
}

func (p *playlistUsecase) AddPlaylist(playlist model.Playlists) error {
	return p.playlist.AddPlaylist(playlist)
}

func (p *playlistUsecase) DeletePlaylist(id string) error {
	return p.playlist.DeletePlaylist(id)
}

func (p *playlistUsecase) GetAllPlaylists() ([]dto.Playlist, error) {
	var playlists []dto.Playlist
	playlistsModel, err := p.playlist.GetAllPlaylists()
	if err != nil {
		return nil, err
	}
	for _, playlistModel := range playlistsModel {
		var playlist dto.Playlist
		dto.TransformPlaylist(&playlistModel, &playlist)
		_, playlist.TotalSong, playlist.TotalDuration, err = p.songUC.GetAllSongsInPlaylist(playlist.ID)
		_, playlist.TotalUser, err = p.userUC.GetAllUsersInPlaylist(playlist.ID)
		playlists = append(playlists, playlist)
	}
	return playlists, nil
}

func (p *playlistUsecase) GetPlaylistByID(id string) (dto.Playlist, error) {
	var playlist dto.Playlist
	playlistModel, err := p.playlist.GetPlaylist(id)
	if err != nil {
		return dto.Playlist{}, err
	}
	dto.TransformPlaylist(&playlistModel, &playlist)
	return playlist, nil
}

func (p *playlistUsecase) GetPlaylistDetail(id string) (dto.PlaylistDetail, error) {
	var playlist dto.PlaylistDetail
	playlistModel, err := p.playlist.GetPlaylist(id)
	if err != nil {
		return dto.PlaylistDetail{}, err
	}
	dto.TransformPlaylistDetail(&playlistModel, &playlist)
	playlist.Song, playlist.TotalSong, playlist.TotalDuration, err = p.songUC.GetAllSongsInPlaylist(id)
	if err != nil {
		return dto.PlaylistDetail{}, err
	}
	playlist.User, playlist.TotalUser, err = p.userUC.GetAllUsersInPlaylist(id)
	if err != nil {
		return dto.PlaylistDetail{}, err
	}
	return playlist, nil
}
