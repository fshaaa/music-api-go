package collab

import (
	"github.com/stretchr/testify/mock"
	"music-api-go/model"
)

type MockCollabRepository struct {
	mock.Mock
}

func NewMockCollabRespository() *MockCollabRepository {
	return &MockCollabRepository{}
}

func (m *MockCollabRepository) GetAllUserID(playlist_id string) ([]string, int, error) {
	ret := m.Called(playlist_id)
	return ret.Get(0).([]string), ret.Get(1).(int), ret.Error(2)
}

func (m *MockCollabRepository) AddCollaboration(collab model.Collaborations) error {
	ret := m.Called(collab)
	return ret.Error(0)
}

func (m *MockCollabRepository) DeleteCollaboration(user_id, playlist_id string) error {
	ret := m.Called(user_id, playlist_id)
	return ret.Error(0)
}
