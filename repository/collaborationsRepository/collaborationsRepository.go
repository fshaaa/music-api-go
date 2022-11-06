package collaborationsRepository

import (
	"database/sql"
	"music-api-go/model"
)

type CollaborationsRepository interface {
	GetAllUserID(playlist_id string) ([]string, error)
	AddCollaboration(collab model.Collaborations) error
	DeleteCollaboration(user_id, playlist_id string) error
}

type collaborationRepository struct {
	db *sql.DB
}

func NewCollaborationRepository(db *sql.DB) *collaborationRepository {
	return &collaborationRepository{db}
}

func (c *collaborationRepository) GetAllUserID(playlist_id string) ([]string, error) {
	var user_id []string
	query := `SELECT user_id FROM collaborations WHERE playlist_id = $1`
	row, err := c.db.Query(query, playlist_id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var id string
		err = row.Scan(&id)
		if err != nil {
			return nil, err
		}
		user_id = append(user_id, id)
	}
	return user_id, nil
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
