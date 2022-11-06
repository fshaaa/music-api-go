package usecase

import (
	"music-api-go/model"
	"music-api-go/repository/albumLikesRepository"
)

type AlbumLikeUsecase interface {
	AddAlbumLike(albumLike model.AlbumLikes) error
	DeleteAlbumLike(user_id, album_id string) error
}

type albumLikeUsecase struct {
	albumLike albumLikesRepository.AlbumLikesRepository
}

func NewAlbumLikeUsecase(al albumLikesRepository.AlbumLikesRepository) *albumLikeUsecase {
	return &albumLikeUsecase{al}
}

func (a *albumLikeUsecase) AddAlbumLike(albumLike model.AlbumLikes) error {
	return a.albumLike.AddAlbumLike(albumLike)
}

func (a *albumLikeUsecase) DeleteAlbumLike(user_id, album_id string) error {
	return a.albumLike.DeleteAlbumLike(user_id, album_id)
}
