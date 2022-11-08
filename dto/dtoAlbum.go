package dto

import "music-api-go/model"

type Album struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Year  string `json:"year"`
	Owner string `json:"owner"`
}

type AlbumDetail struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Year          string `json:"year"`
	Owner         string `json:"owner"`
	TotalLike     int    `json:"total_like"`
	TotalSong     int    `json:"total_song"`
	TotalDuration int    `json:"total_duration"`
	Song          []Song `json:"song"`
}

func TransformAlbum(src *model.Albums, dest *Album) {
	dest.ID = src.ID
	dest.Name = src.Name
	dest.Year = src.Year
	dest.Owner = src.Owner
}

func TransformAlbumDetail(src model.Albums, dest AlbumDetail) {
	dest.ID = src.ID
	dest.Name = src.Name
	dest.Year = src.Year
	dest.Owner = src.Owner
}
