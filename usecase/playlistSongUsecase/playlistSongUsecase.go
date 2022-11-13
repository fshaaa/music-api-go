package playlistSongUsecase

import (
	"music-api-go/model"
	"music-api-go/repository/playlist_song_repository"
)

type PlaylistSongUsecase interface {
	AddPlaylistSong(playlistSong model.PlaylistSongs) error
	DeletePlaylistSong(song_id, playlist_id string) error
}

type playlistSongUsecase struct {
	playlistSong playlist_song_repository.PlaylistSongsRepository
}

func NewPlaylistSongUsecase(ps playlist_song_repository.PlaylistSongsRepository) *playlistSongUsecase {
	return &playlistSongUsecase{ps}
}

func (p *playlistSongUsecase) AddPlaylistSong(playlistSong model.PlaylistSongs) error {
	return p.playlistSong.AddSongInPlaylist(playlistSong)
}

func (p *playlistSongUsecase) DeletePlaylistSong(song_id, playlist_id string) error {
	return p.playlistSong.DeleteSongInPlaylist(song_id, playlist_id)
}
