package albumLikeController

import (
	"github.com/labstack/echo/v4"
	"music-api-go/model"
	"music-api-go/usecase"
	"music-api-go/util/uuid"
	"net/http"
	"time"
)

type AlbumLikeController interface{}

type albumLikeController struct {
	albumLike usecase.AlbumLikeUsecase
}

func NewAlbumLikeController(al usecase.AlbumLikeUsecase) *albumLikeController {
	return &albumLikeController{al}
}

func (a *albumLikeController) AddAlbumLike(c echo.Context) error {
	var albumLike model.AlbumLikes
	c.Bind(&albumLike)
	albumLike.ID = uuid.CreateUUID()
	albumLike.CreatedAt = time.Now().Format(time.RFC1123Z)
	albumLike.UpdatedAt = albumLike.CreatedAt

	err := a.albumLike.AddAlbumLike(albumLike)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "add like album",
	})
}

func (a *albumLikeController) DeleteAlbumLike(c echo.Context) error {
	var albumLikes model.AlbumLikes
	c.Bind(&albumLikes)
	err := a.albumLike.DeleteAlbumLike(albumLikes.User_id, albumLikes.Album_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "delete like album",
	})
}
