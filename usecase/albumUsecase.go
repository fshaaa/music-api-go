package usecase

import (
	"music-api-go/dto"
	"music-api-go/model"
	"music-api-go/repository/albumLikesRepository"
	"music-api-go/repository/albumRepository"
	"music-api-go/repository/songRepository"
)

type AlbumUsecase interface {
	GetAllAlbums() ([]dto.Album, error)
	GetAlbumByID(id string) (dto.Album, error)
	AddAlbum(album model.Albums) error
	DeleteAlbum(id string) error
	GetAlbumDetail(id string) (dto.AlbumDetail, error)
	GetUsersLikeAlbum(id string) ([]dto.User, int, error)
}

type albumUsecase struct {
	album     albumRepository.AlbumRepository
	albumLike albumLikesRepository.AlbumLikesRepository
	song      songRepository.SongRepository
}

func NewAlbumUsecase(a albumRepository.AlbumRepository, al albumLikesRepository.AlbumLikesRepository, s songRepository.SongRepository) *albumUsecase {
	return &albumUsecase{a, al, s}
}

func (a *albumUsecase) GetAllAlbums() ([]dto.Album, error) {
	var albums []dto.Album
	albumsModel, err := a.album.GetAllAlbums()
	if err != nil {
		return nil, err
	}
	for _, albumModel := range albumsModel {
		var album dto.Album
		dto.TransformAlbum(&albumModel, &album)
		albums = append(albums, album)
	}
	return albums, nil
}

func (a *albumUsecase) GetAlbumByID(id string) (dto.Album, error) {
	var album dto.Album
	albumModel, err := a.album.GetAlbum(id)
	if err != nil {
		return dto.Album{}, err
	}
	dto.TransformAlbum(&albumModel, &album)
	return album, nil
}

func (a *albumUsecase) AddAlbum(album model.Albums) error {
	return a.album.AddAlbum(album)
}

func (a *albumUsecase) DeleteAlbum(id string) error {
	return a.album.DeleteAlbum(id)
}

func (a *albumUsecase) GetAlbumDetail(id string) (dto.AlbumDetail, error) {
	var album dto.AlbumDetail
	albumModel, err := a.album.GetAlbum(id)
	dto.TransformAlbumDetail(&albumModel, &album)
	album.TotalLike, err = a.albumLike.GetTotalAlbumLikes(id)
	if err != nil {
		return dto.AlbumDetail{}, err
	}
	album.Song, err = a.song.GetSongsByAlbumID(id)
	if err != nil {
		return dto.AlbumDetail{}, err
	}
	for _, value := range album.Song {
		album.TotalDuration += value.Duration
		album.TotalSong++
	}
	return album, nil
}

func (a *albumUsecase) GetUsersLikeAlbum(id string) ([]dto.User, int, error) {
	users, err := a.albumLike.GetUsersLikeAlbum(id)
	if err != nil {
		return nil, 0, err
	}
	totalUser, err := a.albumLike.GetTotalAlbumLikes(id)
	if err != nil {
		return nil, 0, err
	}
	return users, totalUser, nil
}
