package playlistActivitiesRepository

import (
	"database/sql"
	"music-api-go/model"
)

type PlaylistActivitiesRepository interface {
	AddPlaylistActivity(activity model.PlaylistActivities) error
	DeletePlaylistActivity(id string) error
}

type playlistActivitiesRepository struct {
	db *sql.DB
}

func NewPlaylistActivitiesRepository(db *sql.DB) *playlistActivitiesRepository {
	return &playlistActivitiesRepository{db}
}

func (p *playlistActivitiesRepository) AddPlaylistActivity(activity model.PlaylistActivities) error {
	query := `INSERT INTO playlist_activities VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := p.db.Exec(query, activity.ID, activity.CreatedAt, activity.UpdatedAt, activity.User_id,
		activity.Playlist_id, activity.Action, activity.Time)
	if err != nil {
		return err
	}
	return nil
}

func (p *playlistActivitiesRepository) DeletePlaylistActivity(id string) error {
	query := `DELETE FROM playlist_activities WHERE id = $1`
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
