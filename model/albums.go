package model

type Albums struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	Year      string `json:"year"`
	Owner     string `json:"owner"`
}
