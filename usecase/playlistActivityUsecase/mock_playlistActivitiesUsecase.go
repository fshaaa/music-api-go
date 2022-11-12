package playlistActivityUsecase

import (
	"github.com/stretchr/testify/mock"
	"music-api-go/model"
)

type mockPlaylistActivitiesRepository struct {
	mock.Mock
}

func NewMockPlaylistActivitiesRepository() *mockPlaylistActivitiesRepository {
	return &mockPlaylistActivitiesRepository{}
}

func (m *mockPlaylistActivitiesRepository) AddPlaylistActivity(activity model.PlaylistActivities) error {
	ret := m.Called(activity)
	return ret.Error(0)
}

func (m *mockPlaylistActivitiesRepository) DeletePlaylistActivity(id string) error {
	ret := m.Called(id)
	return ret.Error(0)
}
