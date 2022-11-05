package model

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
