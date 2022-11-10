package dto

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
