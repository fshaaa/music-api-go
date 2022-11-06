package dto

import "music-api-go/model"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type UserToken struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func TransformUser(src model.Users, dest User) {
	dest.ID = src.ID
	dest.Username = src.Username
	dest.Email = src.Email
	dest.Fullname = src.Fullname
}
