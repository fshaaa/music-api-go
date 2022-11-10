package songController

import (
	"github.com/labstack/echo/v4"
	"music-api-go/model"
	"music-api-go/usecase"
	"music-api-go/util/uuid"
	"net/http"
	"time"
)

type SongController interface{}

type songController struct {
	song usecase.SongUsecase
}

func NewSongController(song usecase.SongUsecase) *songController {
	return &songController{song}
}

func (s *songController) GetAllSongs(c echo.Context) error {
	songs, err := s.song.GetAllSongs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get 10 songs",
		"songs":   songs,
	})
}

func (s *songController) GetSongById(c echo.Context) error {
	id := c.Param("id")
	song, err := s.song.GetSongByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get song",
		"song":    song,
	})
}

func (s *songController) AddSong(c echo.Context) error {
	var song model.Songs
	c.Bind(&song)
	song.ID = uuid.CreateUUID()
	song.CreatedAt = time.Now().Format(time.RFC1123Z)
	song.UpdatedAt = song.CreatedAt
	err := s.song.AddSong(song)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success add song",
		"song":    song,
	})
}

func (s *songController) UpdateSong(c echo.Context) error {
	id := c.Param("id")
	var song model.Songs
	c.Bind(&song)
	song.UpdatedAt = time.Now().Format(time.RFC1123Z)
	res, err := s.song.UpdateSong(id, song)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update song",
		"id song": id,
		"update":  res,
	})
}

func (s *songController) DeleteSong(c echo.Context) error {
	id := c.Param("id")
	err := s.song.DeleteSong(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete song",
	})
}

func (s *songController) SearchSong(c echo.Context) error {
	title := c.QueryParam("name")
	songs, err := s.song.SearchSong(title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "search song by tittle",
		"songs":   songs,
	})
}
