package playlistSongsRepository

import (
	"database/sql"
	"music-api-go/model"
)

type PlaylistSongsRepository interface {
	GetAllSongID(playlist_id string) ([]string, error)
	GetDurationPlaylist(playlist_id string) (int, error)
	AddSongInPlaylist(playlist model.PlaylistSongs) error
	DeleteSongInPlaylist(song_id, playlist_id string) error
}

type playlistSongsRepository struct {
	db *sql.DB
}

func NewPlaylistSongsRepository(db *sql.DB) *playlistSongsRepository {
	return &playlistSongsRepository{db}
}

func (p *playlistSongsRepository) GetAllSongID(playlist_id string) ([]string, error) {
	var song_id []string
	query := `SELECT song_id FROM playlist_song WHERE playlist_id = $1`
	row, err := p.db.Query(query, playlist_id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var id string
		err = row.Scan(&id)
		if err != nil {
			return nil, err
		}
		song_id = append(song_id, id)
	}
	return song_id, nil
}

func (p *playlistSongsRepository) GetDurationPlaylist(playlist_id string) (int, error) {
	query := `SELECT SUM(duration) FROM songs s, playlist_songs p WHERE p.playlist_id = $1 AND
				s.id = p.song_id`
	row, err := p.db.Query(query, playlist_id)
	if err != nil {
		return 0, err
	}
	var totalTime = 0
	for row.Next() {
		err = row.Scan(totalTime)
		if err != nil {
			return 0, err
		}
	}
	return totalTime, nil
}

func (p *playlistSongsRepository) AddSongInPlaylist(playlist model.PlaylistSongs) error {
	query := `INSERT INTO playlist_songs VALUES ($1,$2,$3,$4,$5)`
	_, err := p.db.Exec(query, playlist.ID, playlist.CreatedAt, playlist.UpdatedAt, playlist.Song_id,
		playlist.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (p *playlistSongsRepository) DeleteSongInPlaylist(song_id, playlist_id string) error {
	query := `DELETE FROM playlist_songs WHERE song_id = $1 AND playlist_id = $2`
	_, err := p.db.Exec(query, song_id, playlist_id)
	if err != nil {
		return err
	}
	return nil
}
