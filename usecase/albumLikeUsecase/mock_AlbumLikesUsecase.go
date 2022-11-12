package albumLikeUsecase

import (
	"github.com/stretchr/testify/mock"
	"music-api-go/dto"
	"music-api-go/model"
)

type mockAlbumLikesRepository struct {
	mock.Mock
}

func NewMockAlbumLikesRepository() *mockAlbumLikesRepository {
	return &mockAlbumLikesRepository{}
}

func (m *mockAlbumLikesRepository) AddAlbumLike(albumLike model.AlbumLikes) error {
	ret := m.Called(albumLike)
	return ret.Error(0)
}

func (m *mockAlbumLikesRepository) DeleteAlbumLike(user_id, album_id string) error {
	ret := m.Called(user_id, album_id)
	return ret.Error(0)
}

func (m *mockAlbumLikesRepository) GetTotalAlbumLikes(album_id string) (int, error) {
	ret := m.Called(album_id)
	return ret.Get(0).(int), ret.Error(1)
}

func (m *mockAlbumLikesRepository) GetUsersLikeAlbum(album_id string) ([]dto.User, error) {
	ret := m.Called(album_id)
	return ret.Get(0).([]dto.User), ret.Error(1)
}
