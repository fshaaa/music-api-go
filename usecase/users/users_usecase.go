package users

import (
	"music-api-go/dto"
	"music-api-go/model"
	"music-api-go/repository/collab"
	"music-api-go/repository/users"
)

type UserUsecase interface {
	GetAllUsersInPlaylist(playlist_id string) ([]dto.User, int, error)
	CreateUser(user model.Users) error
	LoginUser(user model.Users) (dto.User, error)
	GetUserById(id string) (dto.User, error)
	UpdateUser(id string, user model.Users) (dto.User, error)
	DeleteUser(id string) error
	SearchUser(name string) ([]dto.User, error)
}

type userUsecase struct {
	user   users.UserRepository
	collab collab.CollaborationsRepository
}

func NewUserUsecase(u users.UserRepository, c collab.CollaborationsRepository) *userUsecase {
	return &userUsecase{u, c}
}

func (u *userUsecase) GetAllUsersInPlaylist(playlist_id string) ([]dto.User, int, error) {
	var users []dto.User

	user_id, total, err := u.collab.GetAllUserID(playlist_id)
	if err != nil {
		return nil, 0, err
	}
	for _, id := range user_id {
		userModel, err := u.user.GetUserById(id)
		if err != nil {
			return nil, 0, err
		}
		user := *userModel.ToDTOUser()
		users = append(users, user)
	}
	return users, total, nil
}

func (u *userUsecase) CreateUser(user model.Users) error {
	return u.user.CreateUser(user)
}

func (u *userUsecase) LoginUser(user model.Users) (dto.User, error) {
	userModel, err := u.user.LoginUser(user)
	return *userModel.ToDTOUser(), err
}

func (u *userUsecase) GetUserById(id string) (dto.User, error) {
	userModel, err := u.user.GetUserById(id)
	return *userModel.ToDTOUser(), err
}

func (u *userUsecase) UpdateUser(id string, user model.Users) (dto.User, error) {
	err := u.user.UpdateUser(id, user)
	if err != nil {
		return dto.User{}, err
	}
	return *user.ToDTOUser(), nil
}

func (u *userUsecase) DeleteUser(id string) error {
	return u.user.DeleteUser(id)
}

func (u *userUsecase) SearchUser(name string) ([]dto.User, error) {
	return u.user.SearchUser(name)
}
