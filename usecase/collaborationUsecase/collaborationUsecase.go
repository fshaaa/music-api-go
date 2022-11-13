package collaborationUsecase

import (
	"music-api-go/model"
	"music-api-go/repository/collab"
)

type CollaborationUsecase interface {
	AddCollaboration(collab model.Collaborations) error
	DeleteCollaboration(user_id, playlist_id string) error
}

type collaborationUsecase struct {
	collab collab.CollaborationsRepository
}

func NewCollabUsecase(c collab.CollaborationsRepository) *collaborationUsecase {
	return &collaborationUsecase{c}
}

func (c *collaborationUsecase) AddCollaboration(collab model.Collaborations) error {
	return c.collab.AddCollaboration(collab)
}

func (c *collaborationUsecase) DeleteCollaboration(user_id, playlist_id string) error {
	return c.collab.DeleteCollaboration(user_id, playlist_id)
}
