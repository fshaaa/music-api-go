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
	LoginUser(user model.Users) (dto.User, error)
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

func (u *userUsecase) UpdateUser(id string, user model.Users) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	req := map[string]interface{}{
		"updated_at": user.UpdatedAt,
		"username":   user.Username,
		"email":      user.Email,
		"fullname":   user.Fullname,
	}
	for key, value := range req {
		if value != nil && value != 0 && value != "" {
			res[key] = value
		}
	}
	u.user.UpdateUser(id, res)
	return res, nil
}

func (u *userUsecase) DeleteUser(id string) error {
	return u.user.DeleteUser(id)
}

func (u *userUsecase) SearchUser(name string) ([]dto.User, error) {
	return u.user.SearchUser(name)
}
