package model

import "music-api-go/dto"

type Songs struct {
	ID        string `json:"id" gorm:"primarykey"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Performer string `json:"performer"`
	Genre     string `json:"genre"`
	Duration  int    `json:"duration"`
	Album_id  string `json:"album_Id"`
}

func (s *Songs) ToDTOSong() *dto.Song {
	return &dto.Song{
		ID:        s.ID,
		Title:     s.Title,
		Year:      s.Year,
		Performer: s.Performer,
		Genre:     s.Genre,
		Duration:  s.Duration,
	}
}
