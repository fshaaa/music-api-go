package songController

import (
	"github.com/labstack/echo/v4"
	"music-api-go/model"
	"music-api-go/repository/songRepository"
	"music-api-go/utilities"
	"net/http"
	"time"
)

type SongController interface{}

type songController struct {
	songRepo songRepository.SongRepository
}

func NewSongController(songRepo songRepository.SongRepository) *songController {
	return &songController{songRepo}
}

func (s *songController) GetAllSongs(c echo.Context) error {
	var songs []model.Songs

	songs, err := s.songRepo.GetAllSongs()
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

	song, err := s.songRepo.GetSongById(id)
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
	song.ID = utilities.CreateUUID()
	song.CreatedAt = time.Now().Format(time.RFC1123Z)
	song.UpdatedAt = song.CreatedAt

	err := s.songRepo.AddSong(song)
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

	res, err := s.songRepo.UpdateSong(id, song)
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

	err := s.songRepo.DeleteSong(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete song",
	})
}
