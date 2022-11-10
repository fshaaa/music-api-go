package dto

type Song struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Performer string `json:"performer"`
	Genre     string `json:"genre"`
	Duration  int    `json:"duration"`
}
