package collaborationsRepository

import (
	"github.com/stretchr/testify/mock"
	"music-api-go/model"
)

type mockCollabRepository struct {
	mock.Mock
}

func NewMockCollabRespository() *mockCollabRepository {
	return &mockCollabRepository{}
}

func (m *mockCollabRepository) GetAllUserID(playlist_id string) ([]string, error) {
	ret := m.Called(playlist_id)
	return ret.Get(0).([]string), ret.Error(1)
}

func (m *mockCollabRepository) AddCollaboration(collab model.Collaborations) error {
	ret := m.Called(collab)
	return ret.Error(0)
}

func (m *mockCollabRepository) DeleteCollaboration(user_id, playlist_id string) error {
	ret := m.Called(user_id, playlist_id)
	return ret.Error(0)
}
