package users

import (
	"github.com/stretchr/testify/mock"
	"music-api-go/dto"
	"music-api-go/model"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user model.Users) error {
	ret := m.Called(user)
	return ret.Error(0)
}

func (m *MockUserRepository) LoginUser(user model.Users) (model.Users, error) {
	ret := m.Called(user)
	return ret.Get(0).(model.Users), ret.Error(1)
}

func (m *MockUserRepository) GetUserById(id string) (model.Users, error) {
	ret := m.Called(id)
	return ret.Get(0).(model.Users), ret.Error(1)
}

func (m *MockUserRepository) UpdateUser(id string, user model.Users) error {
	ret := m.Called(id, user)
	return ret.Error(0)
}

func (m *MockUserRepository) DeleteUser(id string) error {
	ret := m.Called(id)
	return ret.Error(0)
}

func (m *MockUserRepository) SearchUser(name string) ([]dto.User, error) {
	ret := m.Called(name)
	return ret.Get(0).([]dto.User), ret.Error(1)
}
