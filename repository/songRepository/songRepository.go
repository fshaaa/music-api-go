package songRepository

import (
	"database/sql"
	"music-api-go/dto"
	"music-api-go/model"
)

type SongRepository interface {
	GetAllSongs() ([]model.Songs, error)
	GetSongById(id string) (model.Songs, error)
	AddSong(song model.Songs) error
	UpdateSong(id, key string, value any) error
	DeleteSong(id string) error
	SearchSong(title string) ([]model.Songs, error)
	GetSongsByAlbumID(id string) ([]dto.Song, error)
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
	return err
}

func (s *songRepository) UpdateSong(id, key string, value any) error {
	query := `UPDATE songs SET $1 = $2 WHERE id = $3`
	_, err := s.db.Exec(query, key, value, id)
	return err
}

func (s *songRepository) DeleteSong(id string) error {
	query := `DELETE FROM songs WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}

func (s *songRepository) SearchSong(title string) ([]model.Songs, error) {
	var songs []model.Songs
	query := `SELECT * FROM songs WHERE title LIKE '%' || $1 || '%' LIMIT 10`
	row, err := s.db.Query(query, title)
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

func (s songRepository) GetSongsByAlbumID(id string) ([]dto.Song, error) {
	var songs []dto.Song
	query := `SELECT * FROM songs WHERE album_id = $1`
	row, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var songDTO dto.Song
		var song model.Songs
		err = row.Scan(&song.ID, &song.CreatedAt, &song.UpdatedAt, &song.Title, &song.Year, &song.Performer,
			&song.Genre, &song.Duration, &song.Album_id)
		if err != nil {
			return nil, err
		}
		dto.TransformSong(&song, &songDTO)
		songs = append(songs, songDTO)
	}
	return songs, nil
}
