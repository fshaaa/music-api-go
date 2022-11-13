package albumRepository

import (
	"github.com/stretchr/testify/mock"
	"music-api-go/model"
)

type mockAlbumRepository struct {
	mock.Mock
}

func NewMockAlbumRepository() *mockAlbumRepository {
	return &mockAlbumRepository{}
}

func (m *mockAlbumRepository) GetAllAlbums() ([]model.Albums, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Albums), ret.Error(1)
}

func (m *mockAlbumRepository) GetAlbum(id string) (model.Albums, error) {
	ret := m.Called(id)
	return ret.Get(0).(model.Albums), ret.Error(1)
}

func (m *mockAlbumRepository) AddAlbum(album model.Albums) error {
	ret := m.Called(album)
	return ret.Error(0)
}

func (m *mockAlbumRepository) DeleteAlbum(id string) error {
	ret := m.Called(id)
	return ret.Error(0)
}
