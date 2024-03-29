package collaborationController

import (
	"github.com/labstack/echo/v4"
	"music-api-go/model"
	"music-api-go/usecase"
	"music-api-go/utilities"
	"net/http"
	"time"
)

type CollabController interface{}

type collabController struct {
	collab usecase.CollaborationUsecase
}

func NewCollabRepository(c usecase.CollaborationUsecase) *collabController {
	return &collabController{c}
}

func (c *collabController) AddCollaboration(e echo.Context) error {
	var collab model.Collaborations
	e.Bind(&collab)
	collab.ID = utilities.CreateUUID()
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
