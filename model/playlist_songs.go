package model

type PlaylistSongs struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Playlist_id string `json:"playlist_id"`
	Song_id     string `json:"song_id"`
}
