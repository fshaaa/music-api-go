package repository

import (
	"music-api-go/repository/albumLikesRepository"
	"music-api-go/repository/albumRepository"
	"music-api-go/repository/collaborationsRepository"
	"music-api-go/repository/playlistActivitiesRepository"
	"music-api-go/repository/playlistSongsRepository"
	"music-api-go/repository/playlistsRepository"
	"music-api-go/repository/songRepository"
	"music-api-go/repository/userRepository"
)

type Repository struct {
	User             userRepository.UserRepository
	Song             songRepository.SongRepository
	Album            albumRepository.AlbumRepository
	AlbumLike        albumLikesRepository.AlbumLikesRepository
	Playlist         playlistsRepository.PlaylistsRepository
	PlaylistSong     playlistSongsRepository.PlaylistSongsRepository
	PlaylistActivity playlistActivitiesRepository.PlaylistActivitiesRepository
	Collaboration    collaborationsRepository.CollaborationsRepository
}
