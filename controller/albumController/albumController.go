package albumController

import (
	"github.com/labstack/echo/v4"
	"music-api-go/model"
	"music-api-go/usecase"
	"music-api-go/utilities"
	"net/http"
	"time"
)

type AlbumController interface{}

type albumController struct {
	album usecase.AlbumUsecase
}

func NewAlbumController(a usecase.AlbumUsecase) *albumController {
	return &albumController{a}
}

func (a *albumController) GetAllAlbums(c echo.Context) error {
	albums, err := a.album.GetAllAlbums()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "get all album",
		"albums":  albums,
	})
}

func (a *albumController) GetAlbumByID(c echo.Context) error {
	id := c.Param("id")
	album, err := a.album.GetAlbumByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "get album by id",
		"album":   album,
	})
}

func (a *albumController) AddAlbum(c echo.Context) error {
	var album model.Albums
	c.Bind(&album)
	album.ID = utilities.CreateUUID()
	album.CreatedAt = time.Now().Format(time.RFC1123Z)
	album.UpdatedAt = album.CreatedAt

	err := a.album.AddAlbum(album)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success add album",
		"album":   album,
	})
}

func (a *albumController) DeleteAlbum(c echo.Context) error {
	id := c.Param("id")
	err := a.album.DeleteAlbum(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete album",
	})
}

func (a *albumController) GetAlbumDetail(c echo.Context) error {
	id := c.Param("id")
	album, err := a.album.GetAlbumDetail(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "get album detail",
		"album":   album,
	})
}

func (a *albumController) GetUsersLikeAlbum(c echo.Context) error {
	id := c.Param("id")
	album, err := a.album.GetAlbumByID(id)
	users, totalLike, err := a.album.GetUsersLikeAlbum(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "get user like album",
		"album":     album.Name,
		"totalLike": totalLike,
		"users":     users,
	})
}
