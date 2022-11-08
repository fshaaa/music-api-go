package dto

import "music-api-go/model"

type Song struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Performer string `json:"performer"`
	Genre     string `json:"genre"`
	Duration  int    `json:"duration"`
}

func TransformSong(src *model.Songs, dest *Song) {
	dest.ID = src.ID
	dest.Title = src.Title
	dest.Year = src.Year
	dest.Performer = src.Performer
	dest.Genre = src.Genre
	dest.Duration = src.Duration
}
