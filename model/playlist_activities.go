package model

type PlaylistActivities struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	User_id     string `json:"user_id"`
	Playlist_id string `json:"playlist_id"`
	Action      string `json:"action"`
	Time        int    `json:"time"`
}
