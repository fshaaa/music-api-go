package userRepository

import (
	"github.com/stretchr/testify/mock"
	"music-api-go/dto"
	"music-api-go/model"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) CreateUser(user model.Users) error {
	ret := m.Called(user)
	return ret.Error(0)
}

func (m *mockUserRepository) LoginUser(user model.Users) (model.Users, error) {
	ret := m.Called(user)
	return ret.Get(0).(model.Users), ret.Error(1)
}

func (m *mockUserRepository) GetUserById(id string) (model.Users, error) {
	ret := m.Called(id)
	return ret.Get(0).(model.Users), ret.Error(1)
}

func (m *mockUserRepository) UpdateUser(id string, req map[string]interface{}) error {
	ret := m.Called(id, req)
	return ret.Error(0)
}

func (m *mockUserRepository) DeleteUser(id string) error {
	ret := m.Called(id)
	return ret.Error(0)
}

func (m *mockUserRepository) SearchUser(name string) ([]dto.User, error) {
	ret := m.Called(name)
	return ret.Get(0).([]dto.User), ret.Error(1)
}
