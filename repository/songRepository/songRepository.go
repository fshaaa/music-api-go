package songRepository

import (
	"database/sql"
	"music-api-go/model"
)

type SongRepository interface {
	GetAllSongs() ([]model.Songs, error)
	GetSongById(id string) (model.Songs, error)
	AddSong(song model.Songs) error
	UpdateSong(id string, song model.Songs) (map[string]interface{}, error)
	DeleteSong(id string) error
}

type songRepository struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *songRepository {
	return &songRepository{db}
}

func (s *songRepository) GetAllSongs() ([]model.Songs, error) {
	var songs []model.Songs
	query := `SELECT * FROM songs LIMIT 10`
	row, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var song model.Songs
		err = row.Scan(&song.ID, &song.CreatedAt, &song.UpdatedAt, &song.Title, &song.Year,
			&song.Performer, &song.Genre, &song.Duration, &song.Album_id)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func (s *songRepository) GetSongById(id string) (model.Songs, error) {
	var song model.Songs
	query := `SELECT * FROM songs WHERE id = $1`
	row, err := s.db.Query(query, id)
	if err != nil {
		return model.Songs{}, err
	}
	for row.Next() {
		err = row.Scan(&song.ID, &song.CreatedAt, &song.UpdatedAt, &song.Title, &song.Year,
			&song.Performer, &song.Genre, &song.Duration, &song.Album_id)
		if err != nil {
			return model.Songs{}, err
		}
	}
	return song, nil
}

func (s *songRepository) AddSong(song model.Songs) error {
	query := `INSERT INTO songs VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	_, err := s.db.Exec(query, song.ID, song.CreatedAt, song.UpdatedAt, song.Title, song.Year,
		song.Performer, song.Genre, song.Duration, song.Album_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *songRepository) UpdateSong(id string, song model.Songs) (map[string]interface{}, error) {
	var res map[string]interface{}
	query := `UPDATE songs SET $1 = $2 WHERE id = $3`
	req := map[string]interface{}{
		"title":     song.Title,
		"year":      song.Year,
		"performer": song.Performer,
		"genre":     song.Genre,
		"duration":  song.Duration,
		"album_id":  song.Album_id,
	}
	for key, value := range req {
		if value != nil {
			_, err := s.db.Exec(query, key, value, id)
			if err != nil {
				return nil, err
			}
			res[key] = value
		}
	}
	return res, nil
}

func (s *songRepository) DeleteSong(id string) error {
	query := `DELETE FROM songs WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
