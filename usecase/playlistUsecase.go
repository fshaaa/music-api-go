package usecase

import (
	"fmt"
	"music-api-go/dto"
	"music-api-go/model"
	"music-api-go/repository/collaborationsRepository"
	"music-api-go/repository/playlistSongsRepository"
	"music-api-go/repository/playlistsRepository"
	"music-api-go/repository/songRepository"
	"music-api-go/repository/userRepository"
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
	collab       collaborationsRepository.CollaborationsRepository
	song         songRepository.SongRepository
	user         userRepository.UserRepository
}

func NewPlaylistUsecase(p playlistsRepository.PlaylistsRepository, ps playlistSongsRepository.PlaylistSongsRepository,
	c collaborationsRepository.CollaborationsRepository, s songRepository.SongRepository, u userRepository.UserRepository) *playlistUsecase {
	return &playlistUsecase{p, ps, c, s, u}
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
		playlist.TotalSong, err = p.playlistSong.GetTotalSongs(playlist.ID)
		playlist.TotalDuration, err = p.playlistSong.GetDurationPlaylist(playlist.ID)
		_, playlist.TotalUserSharing, err = p.collab.GetAllUserID(playlist.ID)
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
	playlist.TotalSong, err = p.playlistSong.GetTotalSongs(id)
	playlist.TotalDuration, err = p.playlistSong.GetDurationPlaylist(id)
	_, playlist.TotalUserSharing, err = p.collab.GetAllUserID(id)
	if err != nil {
		return dto.Playlist{}, err
	}
	return playlist, nil
}

func (p *playlistUsecase) GetPlaylistDetail(id string) (dto.PlaylistDetail, error) {
	var playlist dto.PlaylistDetail
	var users_id, songs_id []string
	var user dto.User
	playlistModel, err := p.playlist.GetPlaylist(id)
	if err != nil {
		return dto.PlaylistDetail{}, err
	}
	dto.TransformPlaylistDetail(&playlistModel, &playlist)
	owner, err := p.user.GetUserById(playlist.Owner)
	dto.TransformUser(&owner, &user)
	playlist.User = append(playlist.User, user)
	songs_id, err = p.playlistSong.GetAllSongID(id)
	for _, s_id := range songs_id {
		var song dto.Song
		songModel, _ := p.song.GetSongById(s_id)
		dto.TransformSong(&songModel, &song)
		playlist.Song = append(playlist.Song, song)
	}
	playlist.TotalSong, err = p.playlistSong.GetTotalSongs(id)
	playlist.TotalDuration, err = p.playlistSong.GetDurationPlaylist(id)
	users_id, playlist.TotalUserSharing, err = p.collab.GetAllUserID(id)
	for _, u_id := range users_id {
		userModel, _ := p.user.GetUserById(u_id)
		dto.TransformUser(&userModel, &user)
		playlist.User = append(playlist.User, user)
	}
	fmt.Println(playlist.Song, playlist.User)
	if err != nil {
		return dto.PlaylistDetail{}, err
	}
	return playlist, nil
}
