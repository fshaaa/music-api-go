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

func TransformUser(src *model.Users, dest *User) {
	dest.ID = src.ID
	dest.Username = src.Username
	dest.Email = src.Email
	dest.Fullname = src.Fullname
}

func CatchValue(src model.Users) []string {
	res := []string{}
	req := map[string]string{
		"updated_at": src.CreatedAt,
		"username":   src.Username,
		"email":      src.Email,
		"fullname":   src.Fullname,
	}
	for _, key := range req {
		res = append(res, key)
	}
	return res
}
