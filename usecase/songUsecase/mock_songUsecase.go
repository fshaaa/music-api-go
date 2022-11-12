package songUsecase

import (
	"github.com/stretchr/testify/mock"
	"music-api-go/dto"
	"music-api-go/model"
)

type mockSongRepository struct {
	mock.Mock
}

func NewMockSongRepository() *mockSongRepository {
	return &mockSongRepository{}
}

func (m *mockSongRepository) GetAllSongs() ([]model.Songs, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Songs), ret.Error(0)
}

func (m *mockSongRepository) GetSongById(id string) (model.Songs, error) {
	ret := m.Called(id)
	return ret.Get(0).(model.Songs), ret.Error(1)
}

func (m *mockSongRepository) AddSong(song model.Songs) error {
	ret := m.Called(song)
	return ret.Error(0)
}

func (m *mockSongRepository) UpdateSong(id string, req map[string]interface{}) error {
	ret := m.Called(id, req)
	return ret.Error(0)
}

func (m *mockSongRepository) DeletSong(id string) error {
	ret := m.Called(id)
	return ret.Error(0)
}

func (m *mockSongRepository) SearchSong(title string) ([]model.Songs, error) {
	ret := m.Called(title)
	return ret.Get(0).([]model.Songs), ret.Error(1)
}

func (m *mockSongRepository) GetSongsByAlbumID(id string) ([]dto.Song, error) {
	ret := m.Called(id)
	return ret.Get(0).([]dto.Song), ret.Error(1)
}
