package playlistSongController

import (
	"github.com/labstack/echo/v4"
	"music-api-go/model"
	"music-api-go/usecase"
	"music-api-go/utilities"
	"net/http"
	"time"
)

type PlaylistSongController interface{}

type playlistSongController struct {
	playlistSong usecase.PlaylistSongUsecase
}

func NewPlaylistSongController(ps usecase.PlaylistSongUsecase) *playlistSongController {
	return &playlistSongController{ps}
}

func (p *playlistSongController) AddPlaylistSong(c echo.Context) error {
	var playlistSong model.PlaylistSongs
	c.Bind(&playlistSong)
	playlistSong.ID = utilities.CreateUUID()
	playlistSong.CreatedAt = time.Now().Format(time.RFC1123Z)
	playlistSong.UpdatedAt = playlistSong.CreatedAt

	err := p.playlistSong.AddPlaylistSong(playlistSong)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "add song in playlist",
	})
}

func (p *playlistSongController) DeletePlaylistSong(c echo.Context) error {
	var playlistSong model.PlaylistSongs
	c.Bind(&playlistSong)

	err := p.playlistSong.DeletePlaylistSong(playlistSong.Song_id, playlistSong.Playlist_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "delete song in playlist",
	})
}
