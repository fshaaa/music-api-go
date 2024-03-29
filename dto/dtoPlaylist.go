package dto

import "music-api-go/model"

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
	User             []User `json:"user"`
	Song             []Song `json:"song"`
}

func TransformPlaylist(src *model.Playlists, dest *Playlist) {
	dest.ID = src.ID
	dest.Name = src.Name
	dest.Owner = src.User_id
}

func TransformPlaylistDetail(src *model.Playlists, dest *PlaylistDetail) {
	dest.ID = src.ID
	dest.Name = src.Name
	dest.Owner = src.User_id
}
