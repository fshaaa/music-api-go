package playlist_song_repository

import (
	"github.com/stretchr/testify/mock"
	"music-api-go/model"
)

type mockPlaylistSongRepository struct {
	mock.Mock
}

func NewMockPlaylistSongRepository() *mockPlaylistSongRepository {
	return &mockPlaylistSongRepository{}
}

func (m *mockPlaylistSongRepository) GetTotalSongs(playlist_id string) (int, error) {
	ret := m.Called(playlist_id)
	return ret.Get(0).(int), ret.Error(1)
}

func (m *mockPlaylistSongRepository) GetAllSongID(playlist_id string) ([]string, error) {
	ret := m.Called(playlist_id)
	return ret.Get(0).([]string), ret.Error(1)
}

func (m *mockPlaylistSongRepository) GetDurationPlaylist(playlist_id string) (int, error) {
	ret := m.Called(playlist_id)
	return ret.Get(0).(int), ret.Error(1)
}

func (m *mockPlaylistSongRepository) AddSongInPlaylist(playlist model.PlaylistSongs) error {
	ret := m.Called(playlist)
	return ret.Error(0)
}

func (m *mockPlaylistSongRepository) DeleteSongInPlaylist(song_id, playlist_id string) error {
	ret := m.Called(song_id, playlist_id)
	return ret.Error(0)
}
