package users

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"music-api-go/model"
	"music-api-go/repository/collab"
	"music-api-go/repository/users"
	"testing"
)

type mockUserUsecase struct {
	mock *users.MockUserRepository
}

func TestCreateUser(t *testing.T) {
	t.Run("success create user", func(t *testing.T) {
		uid := uuid.NewString()

		mockUser := &model.Users{
			ID:       uid,
			Username: "fathazhar",
			Password: "azhar",
			Email:    "fathazhar@gmail.com",
			Fullname: "Fath Azhar",
		}

		mockUserRepository := new(users.MockUserRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		user := NewUserUsecase(mockUserRepository, mockCollabRepository)

		mockUserRepository.On("CreateUser", *mockUser).Return(nil)

		err := user.CreateUser(*mockUser)
		assert.NoError(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestLoginUser(t *testing.T) {
	t.Run("success login user", func(t *testing.T) {
		mockUser := model.Users{
			Username: "fathazhar",
			Password: "azhar",
		}

		mockUserRepository := new(users.MockUserRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		user := NewUserUsecase(mockUserRepository, mockCollabRepository)

		mockUserRepository.On("LoginUser", mock.Anything).Return(mockUser, nil)

		users, err := user.LoginUser(mockUser)
		assert.NoError(t, err)
		assert.Error(t, users, mockUser)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestGetUserById(t *testing.T) {
	t.Run("success Get User By ID", func(t *testing.T) {
		uid := uuid.NewString()

		mockUser := model.Users{
			ID:       uid,
			Username: "fathazhar",
			Email:    "fathazhar@gmail.com",
			Fullname: "Fath Azhar",
		}

		mockUserRepository := new(users.MockUserRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		user := NewUserUsecase(mockUserRepository, mockCollabRepository)

		mockUserRepository.On("GetUserById", uid).Return(mockUser, nil)

		users, err := user.GetUserById(uid)
		assert.NoError(t, err)
		assert.Error(t, users, mockUser)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("success Update User", func(t *testing.T) {
		uid := uuid.NewString()

		mockUser := model.Users{
			ID:       uid,
			Username: "fathazhar",
			Email:    "fathazhar@gmail.com",
			Fullname: "Fath Azhar",
		}

		mockUserRepository := new(users.MockUserRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		user := NewUserUsecase(mockUserRepository, mockCollabRepository)

		mockUserRepository.On("UpdateUser", uid, mockUser).Return(*mockUser.ToDTOUser(), nil)

		users, err := user.UpdateUser(uid, mockUser)
		assert.NoError(t, err)
		assert.Error(t, users, mockUser)
		mockUserRepository.AssertExpectations(t)
	})
}

//
//func (m *mockUserUsecase) LoginUser(user model.Users) (dto.User, error) {
//	ret := m.Called(user)
//	return ret.Get(0).(dto.User), ret.Error(1)
//}
//
//func (m *mockUserUsecase) GetUserById(id string) (dto.User, error) {
//	ret := m.Called(id)
//	return ret.Get(0).(dto.User), ret.Error(1)
//}
//
//func (m *mockUserUsecase) UpdateUser(id string, user model.Users) (map[string]interface{}, error) {
//	ret := m.Called(id, user)
//	return ret.Get(0).(map[string]interface{}), ret.Error(1)
//}
//
//func (m *mockUserUsecase) DeleteUser(id string) error {
//	ret := m.Called(id)
//	return ret.Error(0)
//}
//
//func (m *mockUserUsecase) SearchUser(name string) ([]dto.User, error) {
//	ret := m.Called(name)
//	return ret.Get(0).([]dto.User), ret.Error(1)
//}
