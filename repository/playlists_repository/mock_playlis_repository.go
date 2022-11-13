package playlists_repository

import (
	"github.com/stretchr/testify/mock"
	"music-api-go/model"
)

type mockPlaylistsRepository struct {
	mock.Mock
}

func (m *mockPlaylistsRepository) GetAllPlaylists() ([]model.Playlists, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Playlists), ret.Error(1)
}

func (m *mockPlaylistsRepository) GetPlaylist(id string) (model.Playlists, error) {
	ret := m.Called(id)
	return ret.Get(0).(model.Playlists), ret.Error(1)
}

func (m *mockPlaylistsRepository) AddPlaylist(playlist model.Playlists) error {
	ret := m.Called(playlist)
	return ret.Error(0)
}

func (m *mockPlaylistsRepository) DeletePlaylist(id string) error {
	ret := m.Called(id)
	return ret.Error(0)
}
