package usecase

import (
	"music-api-go/dto"
	"music-api-go/model"
	"music-api-go/repository/collaborationsRepository"
	"music-api-go/repository/userRepository"
)

type UserUsecase interface {
	GetAllUsersInPlaylist(playlist_id string) ([]dto.User, int, error)
	CreateUser(user model.Users) error
	LoginUser(user model.Users) (dto.UserToken, error)
	GetUserById(id string) (dto.User, error)
	UpdateUser(id string, user model.Users) (map[string]interface{}, error)
	DeleteUser(id string) error
	SearchUser(name string) ([]dto.User, error)
}

type userUsecase struct {
	user   userRepository.UserRepository
	collab collaborationsRepository.CollaborationsRepository
}

func NewUserUsecase(u userRepository.UserRepository, c collaborationsRepository.CollaborationsRepository) *userUsecase {
	return &userUsecase{u, c}
}

func (u *userUsecase) GetAllUsersInPlaylist(playlist_id string) ([]dto.User, int, error) {
	var users []dto.User
	var totalUser = 0

	user_id, err := u.collab.GetAllUserID(playlist_id)
	if err != nil {
		return nil, 0, err
	}
	for _, id := range user_id {
		var user dto.User
		userModel, err := u.user.GetUserById(id)
		if err != nil {
			return nil, 0, err
		}
		dto.TransformUser(userModel, user)
		users = append(users, user)
		totalUser++
	}
	return users, totalUser, nil
}

func (u *userUsecase) CreateUser(user model.Users) error {
	return u.user.CreateUser(user)
}

func (u *userUsecase) LoginUser(user model.Users) (dto.UserToken, error) {
	var DTOuser dto.User
	userModel, err := u.user.LoginUser(user)
	dto.TransformUser(userModel, DTOuser)
	return DTOuser, err
}

func (u *userUsecase) GetUserById(id string) (dto.User, error) {
	var DTOuser dto.User
	userModel, err := u.user.GetUserById(id)
	dto.TransformUser(userModel, DTOuser)
	return DTOuser, err
}

func (u *userUsecase) UpdateUser(id string, user model.Users) (map[string]interface{}, error) {
	var update map[string]interface{}
	req := map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
		"fullname": user.Fullname,
	}
	var err error
	for key, value := range req {
		if value != nil {
			err = u.user.UpdateUser(id, key, value)
			update[key] = value
		}
	}
	return update, err
}

func (u *userUsecase) DeleteUser(id string) error {
	return u.user.DeleteUser(id)
}

func (u *userUsecase) SearchUser(name string) ([]dto.User, error) {
	return u.user.SearchUser(name)
}
