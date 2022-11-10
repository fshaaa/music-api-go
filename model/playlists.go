package model

import "music-api-go/dto"

type Playlists struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	User_id   string `json:"user_id"`
}

func (p *Playlists) ToDTOPlaylists() *dto.Playlist {
	return &dto.Playlist{
		ID:    p.ID,
		Name:  p.Name,
		Owner: p.User_id,
	}
}

func (p Playlists) ToDTOPlaylistDetails() *dto.PlaylistDetail {
	return &dto.PlaylistDetail{
		ID:    p.ID,
		Name:  p.Name,
		Owner: p.User_id,
	}
}
