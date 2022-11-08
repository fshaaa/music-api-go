package playlistController

import (
	"github.com/labstack/echo/v4"
	"music-api-go/model"
	"music-api-go/usecase"
	"music-api-go/utilities"
	"net/http"
	"time"
)

type PlaylistController interface{}

type playlistController struct {
	playlist usecase.PlaylistUsecase
}

func NewPlaylistController(p usecase.PlaylistUsecase) *playlistController {
	return &playlistController{p}
}

func (p *playlistController) GetAllPlaylists(c echo.Context) error {
	playlists, err := p.playlist.GetAllPlaylists()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "get all playlists 10 row",
		"playlists": playlists,
	})
}

func (p *playlistController) GetPlaylitsDetail(c echo.Context) error {
	id := c.Param("id")
	playlist, err := p.playlist.GetPlaylistDetail(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":  "get detail playlist",
		"playlist": playlist,
	})
}

func (p *playlistController) GetPlaylistById(c echo.Context) error {
	id := c.Param("id")
	playlist, err := p.playlist.GetPlaylistByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":  "get playlist by id",
		"playlist": playlist,
	})
}

func (p *playlistController) AddPlaylist(c echo.Context) error {
	var playlist model.Playlists
	c.Bind(&playlist)
	playlist.ID = utilities.CreateUUID()
	playlist.CreatedAt = time.Now().Format(time.RFC1123Z)
	playlist.UpdatedAt = playlist.CreatedAt

	err := p.playlist.AddPlaylist(playlist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":  "success add playlist",
		"playlist": playlist,
	})
}

func (p *playlistController) DeletePlaylist(c echo.Context) error {
	id := c.Param("id")
	err := p.playlist.DeletePlaylist(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete playlist",
	})
}
