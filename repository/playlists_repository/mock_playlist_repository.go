package playlists_repository

import (
	"github.com/stretchr/testify/mock"
	"music-api-go/model"
)

type MockPlaylistsRepository struct {
	mock.Mock
}

func (m *MockPlaylistsRepository) GetAllPlaylists() ([]model.Playlists, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Playlists), ret.Error(1)
}

func (m *MockPlaylistsRepository) GetPlaylist(id string) (model.Playlists, error) {
	ret := m.Called(id)
	return ret.Get(0).(model.Playlists), ret.Error(1)
}

func (m *MockPlaylistsRepository) AddPlaylist(playlist model.Playlists) error {
	ret := m.Called(playlist)
	return ret.Error(0)
}

func (m *MockPlaylistsRepository) DeletePlaylist(id string) error {
	ret := m.Called(id)
	return ret.Error(0)
}

func (m *MockPlaylistsRepository) GetPlaylistByUser(user_id string) ([]model.Playlists, error) {
	ret := m.Called(user_id)
	return ret.Get(0).([]model.Playlists), ret.Error(1)
}
