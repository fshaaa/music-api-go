package dto

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
