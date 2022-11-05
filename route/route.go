package route

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"music-api-go/config"
	"music-api-go/controller/songController"
	userC "music-api-go/controller/userController"
	songR "music-api-go/repository/songRepository"
	userR "music-api-go/repository/userRepository"
)

func NewRoute(e *echo.Echo, db *sql.DB) {
	userRepos := userR.NewUserRepository(db)
	userControl := userC.NewUserController(userRepos)

	appUser := e.Group("")
	appUser.POST("/signup", userControl.CreateUser)
	appUser.POST("/login", userControl.LoginUser)

	appUserJWT := e.Group("/users")
	appUserJWT.Use(middleware.JWT([]byte(config.Cfg.TokenSecret)))
	appUserJWT.GET("/:id", userControl.GetUserById)
	appUserJWT.PUT("/:id", userControl.UpdateUser)
	appUserJWT.DELETE("/:id", userControl.DeleteUser)
	appUserJWT.GET("/search", userControl.SearchUser)

	songRepos := songR.NewSongRepository(db)
	songControl := songController.NewSongController(songRepos)

	appSong := e.Group("/songs")
	appSong.Use(middleware.JWT([]byte(config.Cfg.TokenSecret)))
	appSong.GET("", songControl.GetAllSongs)
	appSong.GET("/:id", songControl.GetSongById)
	appSong.POST("", songControl.AddSong)
	appSong.PUT("/:id", songControl.UpdateSong)
	appSong.DELETE("/:id", userControl.DeleteUser)
}
