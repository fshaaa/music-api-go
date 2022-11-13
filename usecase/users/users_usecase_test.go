package users

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"music-api-go/dto"
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
			Username: "fathazhar",
			Email:    "fathazhar@gmail.com",
			Fullname: "Fath Azhar",
		}

		mockUserRepository := new(users.MockUserRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		user := NewUserUsecase(mockUserRepository, mockCollabRepository)

		mockUserRepository.On("UpdateUser", uid, mockUser).Return(nil)

		users, err := user.UpdateUser(uid, mockUser)
		assert.NoError(t, err)
		assert.Error(t, users, mockUser)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("success Delete User", func(t *testing.T) {
		uid := uuid.NewString()

		mockUserRepository := new(users.MockUserRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		user := NewUserUsecase(mockUserRepository, mockCollabRepository)

		mockUserRepository.On("DeleteUser", uid).Return(nil)

		err := user.DeleteUser(uid)
		assert.NoError(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestSearchUser(t *testing.T) {
	t.Run("success Search User", func(t *testing.T) {
		var res []dto.User

		mockUser := model.Users{
			Username: "fathazhar",
			Email:    "fathazhar@gmail.com",
			Fullname: "Fath Azhar",
		}

		mockUserRepository := new(users.MockUserRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		user := NewUserUsecase(mockUserRepository, mockCollabRepository)

		mockUserRepository.On("SearchUser", mock.Anything).Return(res, nil)

		_, err := user.SearchUser(mockUser.Fullname)
		assert.NoError(t, err)
		//assert.Error(t, users, res)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestGetAllUsersInPlaylist(t *testing.T) {
	t.Run("success Get All Users in Playlist", func(t *testing.T) {
		uid := uuid.NewString()
		res := model.Users{
			ID:       uuid.NewString(),
			Username: "fathazhar",
			Email:    "fathazhar@gmail.com",
			Fullname: "Fath Azhar",
		}
		total := 1
		arr := []string{
			uuid.NewString(),
		}

		mockUserRepository := new(users.MockUserRepository)
		mockCollabRepository := new(collab.MockCollabRepository)
		user := NewUserUsecase(mockUserRepository, mockCollabRepository)

		mockCollabRepository.On("GetAllUserID", uid).Return(arr, total, nil)
		mockUserRepository.On("GetUserById", mock.Anything).Return(res, nil)

		_, _, err := user.GetAllUsersInPlaylist(uid)
		assert.NoError(t, err)
		//assert.Error(t, tot, total)
		mockUserRepository.AssertExpectations(t)
	})
}
