package collaborationController

import (
	"github.com/labstack/echo/v4"
	"music-api-go/model"
	"music-api-go/usecase/collaborationUsecase"
	"music-api-go/util/uuid"
	"net/http"
	"time"
)

type CollabController interface{}

type collabController struct {
	collab collaborationUsecase.CollaborationUsecase
}

func NewCollabRepository(c collaborationUsecase.CollaborationUsecase) *collabController {
	return &collabController{c}
}

func (c *collabController) AddCollaboration(e echo.Context) error {
	var collab model.Collaborations
	e.Bind(&collab)
	collab.ID = uuid.CreateUUID()
	collab.CreatedAt = time.Now().Format(time.RFC1123Z)
	collab.UpdatedAt = collab.CreatedAt

	err := c.collab.AddCollaboration(collab)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, echo.Map{
		"message": "success add collaboration",
		"collab":  collab,
	})
}

func (c *collabController) DeleteCollaboration(e echo.Context) error {
	var collab model.Collaborations
	e.Bind(&collab)
	err := c.collab.DeleteCollaboration(collab.User_id, collab.Playlist_id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, echo.Map{
		"message": "success delete collaboration",
	})
}
