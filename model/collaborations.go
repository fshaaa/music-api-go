package model

type Collaborations struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	User_id     string `json:"user_id"`
	Playlist_id string `json:"playlist_id"`
}
