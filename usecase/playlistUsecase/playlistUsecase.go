package playlistUsecase

import (
	"music-api-go/dto"
	"music-api-go/model"
	"music-api-go/repository/collab"
	"music-api-go/repository/playlist_song_repository"
	"music-api-go/repository/playlists_repository"
	"music-api-go/repository/song_repository"
	"music-api-go/repository/users"
)

type PlaylistUsecase interface {
	AddPlaylist(playlist model.Playlists) error
	DeletePlaylist(id string) error
	GetPlaylistByID(id string) (dto.Playlist, error)
	GetAllPlaylists() ([]dto.Playlist, error)
	GetPlaylistDetail(id string) (dto.PlaylistDetail, error)
	GetPlaylistByUser(id string) ([]dto.Playlist, error)
}

type playlistUsecase struct {
	playlist     playlists_repository.PlaylistsRepository
	playlistSong playlist_song_repository.PlaylistSongsRepository
	collab       collab.CollaborationsRepository
	song         song_repository.SongRepository
	user         users.UserRepository
}

func NewPlaylistUsecase(p playlists_repository.PlaylistsRepository, ps playlist_song_repository.PlaylistSongsRepository,
	c collab.CollaborationsRepository, s song_repository.SongRepository, u users.UserRepository) *playlistUsecase {
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
		playlist := *playlistModel.ToDTOPlaylists()
		playlist.TotalSong, err = p.playlistSong.GetTotalSongs(playlist.ID)
		playlist.TotalDuration, err = p.playlistSong.GetDurationPlaylist(playlist.ID)
		_, playlist.TotalUserSharing, err = p.collab.GetAllUserID(playlist.ID)
		playlists = append(playlists, playlist)
	}
	return playlists, nil
}

func (p *playlistUsecase) GetPlaylistByID(id string) (dto.Playlist, error) {
	playlistModel, err := p.playlist.GetPlaylist(id)
	if err != nil {
		return dto.Playlist{}, err
	}
	playlist := *playlistModel.ToDTOPlaylists()
	playlist.TotalSong, err = p.playlistSong.GetTotalSongs(id)
	playlist.TotalDuration, err = p.playlistSong.GetDurationPlaylist(id)
	_, playlist.TotalUserSharing, err = p.collab.GetAllUserID(id)
	if err != nil {
		return dto.Playlist{}, err
	}
	return playlist, nil
}

func (p *playlistUsecase) GetPlaylistDetail(id string) (dto.PlaylistDetail, error) {
	var users_id, songs_id []string
	playlistModel, err := p.playlist.GetPlaylist(id)
	if err != nil {
		return dto.PlaylistDetail{}, err
	}
	playlist := *playlistModel.ToDTOPlaylistDetails()
	owner, err := p.user.GetUserById(playlist.Owner)
	user := *owner.ToDTOUser()
	playlist.User = append(playlist.User, user)
	songs_id, err = p.playlistSong.GetAllSongID(id)
	for _, s_id := range songs_id {
		songModel, _ := p.song.GetSongById(s_id)
		song := *songModel.ToDTOSong()
		playlist.Song = append(playlist.Song, song)
	}
	playlist.TotalSong, err = p.playlistSong.GetTotalSongs(id)
	playlist.TotalDuration, err = p.playlistSong.GetDurationPlaylist(id)
	users_id, playlist.TotalUserSharing, err = p.collab.GetAllUserID(id)
	for _, u_id := range users_id {
		userModel, _ := p.user.GetUserById(u_id)
		user = *userModel.ToDTOUser()
		playlist.User = append(playlist.User, user)
	}
	if err != nil {
		return dto.PlaylistDetail{}, err
	}
	return playlist, nil
}

func (p *playlistUsecase) GetPlaylistByUser(user_id string) ([]dto.Playlist, error) {
	var playlists []dto.Playlist
	playlistsModel, err := p.playlist.GetPlaylistByUser(user_id)
	if err != nil {
		return nil, err
	}
	for _, playlistModel := range playlistsModel {
		playlists = append(playlists, *playlistModel.ToDTOPlaylists())
	}
	return playlists, nil
}
