package collaborationsRepository

import (
	"database/sql"
	"music-api-go/model"
)

type CollaborationsRepository interface {
	AddCollaboration(collab model.Collaborations) error
	DeleteCollaboration(user_id, playlist_id string) error
}

type collaborationRepository struct {
	db *sql.DB
}

func NewCollaborationRepository(db *sql.DB) *collaborationRepository {
	return &collaborationRepository{db}
}

func (c *collaborationRepository) AddCollaboration(collab model.Collaborations) error {
	query := `INSERT INTO collaborations VALUES ($1,$2,$3,$4,$5)`
	_, err := c.db.Exec(query, collab.ID, collab.CreatedAt, collab.UpdatedAt, collab.User_id,
		collab.Playlist_id)
	if err != nil {
		return err
	}
	return nil
}

func (c *collaborationRepository) DeleteCollaboration(user_id, playlist_id string) error {
	query := `DELETE FROM collaborations WHERE user_id = $1 AND playlist_id = $2`
	_, err := c.db.Exec(query, user_id, playlist_id)
	if err != nil {
		return err
	}
	return nil
}
