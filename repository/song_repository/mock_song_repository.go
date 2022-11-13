package song_repository

import (
	"github.com/stretchr/testify/mock"
	"music-api-go/dto"
	"music-api-go/model"
)

type MockSongRepository struct {
	mock.Mock
}

func NewMockSongRepository() *MockSongRepository {
	return &MockSongRepository{}
}

func (m *MockSongRepository) GetAllSongs() ([]model.Songs, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Songs), ret.Error(1)
}

func (m *MockSongRepository) GetSongById(id string) (model.Songs, error) {
	ret := m.Called(id)
	return ret.Get(0).(model.Songs), ret.Error(1)
}

func (m *MockSongRepository) AddSong(song model.Songs) error {
	ret := m.Called(song)
	return ret.Error(0)
}

func (m *MockSongRepository) UpdateSong(id, key string, value any) error {
	ret := m.Called(id, key, value)
	return ret.Error(0)
}

func (m *MockSongRepository) DeleteSong(id string) error {
	ret := m.Called(id)
	return ret.Error(0)
}

func (m *MockSongRepository) SearchSong(title string) ([]model.Songs, error) {
	ret := m.Called(title)
	return ret.Get(0).([]model.Songs), ret.Error(1)
}

func (m *MockSongRepository) GetSongsByAlbumID(id string) ([]dto.Song, error) {
	ret := m.Called(id)
	return ret.Get(0).([]dto.Song), ret.Error(1)
}
