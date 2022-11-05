package model

type AlbumLikes struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	User_id   string `json:"user_id"`
	Album_id  string `json:"album_id"`
}
