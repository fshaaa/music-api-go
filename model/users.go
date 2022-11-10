package model

import "music-api-go/dto"

type Users struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
}

func (u *Users) ToDTOUser() *dto.User {
	return &dto.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Fullname: u.Fullname,
	}
}
