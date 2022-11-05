package playlistsRepository

import (
	"database/sql"
	"music-api-go/model"
)

type PlaylistsRepository interface {
	GetAllPlaylists() ([]model.Playlists, error)
	GetPlaylist(id string) (model.Playlists, error)
	AddPlaylist(playlist model.Playlists) error
	DeletePlaylist(id string) error
}

type playlistsRepository struct {
	db *sql.DB
}

func NewPlaylistRepository(db *sql.DB) *playlistsRepository {
	return &playlistsRepository{db}
}

func (p *playlistsRepository) GetAllPlaylists() ([]model.Playlists, error) {
	var playlists []model.Playlists
	query := `SELECT * FROM playlists LIMIT 10`

	row, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var playlist model.Playlists
		err = row.Scan(&playlist.ID, &playlist.CreatedAt, &playlist.UpdatedAt, &playlist.Name, &playlist.User_id)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func (p *playlistsRepository) GetPlaylist(id string) (model.Playlists, error) {
	var playlist model.Playlists
	query := `SELECT * FROM playlists WHERE id = ?`

	row, err := p.db.Query(query, id)
	if err != nil {
		return model.Playlists{}, err
	}
	for row.Next() {
		err = row.Scan(&playlist.ID, &playlist.CreatedAt, &playlist.UpdatedAt, &playlist.Name,
			&playlist.User_id)
	}

	return playlist, nil
}

func (p playlistsRepository) AddPlaylist(playlist model.Playlists) error {
	query := `INSERT INTO playlists VALUES($1,$2,$3,$4,$5)`
	_, err := p.db.Exec(query, playlist.ID, playlist.CreatedAt, playlist.UpdatedAt, playlist.Name, playlist.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *playlistsRepository) DeletePlaylist(id string) error {
	query := `DELETE FROM playlists WHERE id = $1`
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
