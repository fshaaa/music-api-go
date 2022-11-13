package playlist_song_repository

import (
	"github.com/stretchr/testify/mock"
	"music-api-go/model"
)

type MockPlaylistSongRepository struct {
	mock.Mock
}

func NewMockPlaylistSongRepository() *MockPlaylistSongRepository {
	return &MockPlaylistSongRepository{}
}

func (m *MockPlaylistSongRepository) GetTotalSongs(playlist_id string) (int, error) {
	ret := m.Called(playlist_id)
	return ret.Get(0).(int), ret.Error(1)
}

func (m *MockPlaylistSongRepository) GetAllSongID(playlist_id string) ([]string, error) {
	ret := m.Called(playlist_id)
	return ret.Get(0).([]string), ret.Error(1)
}

func (m *MockPlaylistSongRepository) GetDurationPlaylist(playlist_id string) (int, error) {
	ret := m.Called(playlist_id)
	return ret.Get(0).(int), ret.Error(1)
}

func (m *MockPlaylistSongRepository) AddSongInPlaylist(playlist model.PlaylistSongs) error {
	ret := m.Called(playlist)
	return ret.Error(0)
}

func (m *MockPlaylistSongRepository) DeleteSongInPlaylist(song_id, playlist_id string) error {
	ret := m.Called(song_id, playlist_id)
	return ret.Error(0)
}
