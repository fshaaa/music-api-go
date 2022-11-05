package playlistSongsRepository

import (
	"database/sql"
	"music-api-go/model"
)

type PlaylistSongsRepository interface {
	AddSongInPlaylist(playlist model.PlaylistSongs) error
	DeleteSongInPlaylist(playlist model.PlaylistSongs) error
}

type playlistSongsRepository struct {
	db *sql.DB
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
