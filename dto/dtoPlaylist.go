package dto

type Playlist struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Owner            string `json:"owner"`
	TotalSong        int    `json:"total_song"`
	TotalDuration    int    `json:"total_duration"`
	TotalUserSharing int    `json:"total_user_sharing"`
}

type PlaylistDetail struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Owner            string `json:"owner"`
	TotalSong        int    `json:"total_song"`
	TotalDuration    int    `json:"total_duration"`
	TotalUserSharing int    `json:"total_user_sharing"`
	User             []User `json:"users"`
	Song             []Song `json:"song"`
}
