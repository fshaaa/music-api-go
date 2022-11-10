package model

import "music-api-go/dto"

type Albums struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	Year      string `json:"year"`
	Owner     string `json:"owner"`
}

func (a *Albums) ToDTOAlbum() *dto.Album {
	return &dto.Album{
		ID:    a.ID,
		Name:  a.Name,
		Year:  a.Year,
		Owner: a.Owner,
	}
}

func (a *Albums) ToDTOAlbumDetail() *dto.AlbumDetail {
	return &dto.AlbumDetail{
		ID:    a.ID,
		Name:  a.Name,
		Year:  a.Year,
		Owner: a.Owner,
	}
}
