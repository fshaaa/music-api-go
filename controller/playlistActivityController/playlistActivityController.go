package playlistActivityController

import (
	"github.com/labstack/echo/v4"
	"music-api-go/model"
	"music-api-go/usecase"
	"music-api-go/utilities"
	"net/http"
	"time"
)

type PlaylistActivityController interface{}

type playlistActivityController struct {
	playlistActivity usecase.PlaylistActivityUsecase
}

func NewPlaylistActivityController(pa usecase.PlaylistActivityUsecase) *playlistActivityController {
	return &playlistActivityController{pa}
}

func (p *playlistActivityController) AddPlaylistActivity(c echo.Context) error {
	var playlistActivity model.PlaylistActivities
	c.Bind(&playlistActivity)
	playlistActivity.ID = utilities.CreateUUID()
	playlistActivity.CreatedAt = time.Now().Format(time.RFC1123Z)
	playlistActivity.UpdatedAt = playlistActivity.CreatedAt

	err := p.playlistActivity.AddPlaylistActivity(playlistActivity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "add playlist activity",
	})
}

func (p *playlistActivityController) DeletePlaylistActivity(c echo.Context) error {
	id := c.Param("id")
	err := p.playlistActivity.DeletePlaylistActivity(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "delete playlist activity",
	})
}
