package route

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"music-api-go/config"
	"music-api-go/controller/albumController"
	"music-api-go/controller/albumLikeController"
	"music-api-go/controller/collaborationController"
	"music-api-go/controller/playlistActivityController"
	"music-api-go/controller/playlistController"
	"music-api-go/controller/playlistSongController"
	"music-api-go/controller/songController"
	"music-api-go/controller/userController"
	"music-api-go/repository/albumLikesRepository"
	"music-api-go/repository/albumRepository"
	"music-api-go/repository/collaborationsRepository"
	"music-api-go/repository/playlistActivitiesRepository"
	"music-api-go/repository/playlistSongsRepository"
	"music-api-go/repository/playlistsRepository"
	"music-api-go/repository/songRepository"
	"music-api-go/repository/userRepository"
	"music-api-go/usecase"
)

func NewRoute(e *echo.Echo, db *sql.DB) {

	albumLikeRepo := albumLikesRepository.NewAlbumLikesRepository(db)
	albumRepo := albumRepository.NewAlbumRepository(db)
	collabRepo := collaborationsRepository.NewCollaborationRepository(db)
	playlistRepo := playlistsRepository.NewPlaylistRepository(db)
	playSongRepo := playlistSongsRepository.NewPlaylistSongsRepository(db)
	playActiv := playlistActivitiesRepository.NewPlaylistActivitiesRepository(db)
	songRepo := songRepository.NewSongRepository(db)
	userRepo := userRepository.NewUserRepository(db)

	albumUc := usecase.NewAlbumUsecase(albumRepo, albumLikeRepo, songRepo)
	albumlikeUc := usecase.NewAlbumLikeUsecase(albumLikeRepo)
	collabUc := usecase.NewCollabUsecase(collabRepo)
	playlistUc := usecase.NewPlaylistUsecase(playlistRepo, playSongRepo, collabRepo, songRepo, userRepo)
	playSongUc := usecase.NewPlaylistSongUsecase(playSongRepo)
	playActivUc := usecase.PlaylistActivityUsecase(playActiv)
	songUc := usecase.NewSongUsecase(songRepo, playSongRepo)
	userUc := usecase.NewUserUsecase(userRepo, collabRepo)

	albumControl := albumController.NewAlbumController(albumUc)
	albumLikeControl := albumLikeController.NewAlbumLikeController(albumlikeUc)
	collabControl := collaborationController.NewCollabRepository(collabUc)
	playlistCotrol := playlistController.NewPlaylistController(playlistUc)
	playSongControl := playlistSongController.NewPlaylistSongController(playSongUc)
	playActivControl := playlistActivityController.NewPlaylistActivityController(playActivUc)
	songControl := songController.NewSongController(songUc)
	userControl := userController.NewUserController(userUc)

	e.POST("/signup", userControl.CreateUser)
	e.POST("/login", userControl.LoginUser)

	appUserJWT := e.Group("/users")
	appUserJWT.Use(middleware.JWT([]byte(config.Cfg.TokenSecret)))
	appUserJWT.GET("/:id", userControl.GetUserById)
	appUserJWT.PUT("/:id", userControl.UpdateUser)
	appUserJWT.DELETE("/:id", userControl.DeleteUser)
	appUserJWT.GET("/search", userControl.SearchUser)

	appSong := e.Group("/songs")
	appSong.Use(middleware.JWT([]byte(config.Cfg.TokenSecret)))
	appSong.GET("", songControl.GetAllSongs)
	appSong.GET("/:id", songControl.GetSongById)
	appSong.POST("", songControl.AddSong)
	appSong.PUT("/:id", songControl.UpdateSong)
	appSong.DELETE("/:id", songControl.DeleteSong)
	appSong.GET("/search", songControl.SearchSong)

	appPlaylist := e.Group("/playlists")
	appPlaylist.Use(middleware.JWT([]byte(config.Cfg.TokenSecret)))
	appPlaylist.GET("", playlistCotrol.GetAllPlaylists)
	appPlaylist.GET("/:id", playlistCotrol.GetPlaylistById)
	appPlaylist.GET("/details/:id", playlistCotrol.GetPlaylitsDetail)
	appPlaylist.POST("", playlistCotrol.AddPlaylist)
	appPlaylist.DELETE("/:id", playlistCotrol.DeletePlaylist)

	appAlbum := e.Group("/albums")
	appAlbum.Use(middleware.JWT([]byte(config.Cfg.TokenSecret)))
	appAlbum.GET("", albumControl.GetAllAlbums)
	appAlbum.GET("/:id", albumControl.GetAlbumByID)
	appAlbum.GET("/details/:id", albumControl.GetAlbumDetail)
	appAlbum.GET("/like/:id", albumControl.GetUsersLikeAlbum)
	appAlbum.POST("", albumControl.AddAlbum)
	appAlbum.DELETE("", albumControl.DeleteAlbum)

	appPlaylistSong := appPlaylist.Group("/songs")
	appPlaylistSong.Use(middleware.JWT([]byte(config.Cfg.TokenSecret)))
	appPlaylistSong.POST("", playSongControl.AddPlaylistSong)
	appPlaylistSong.DELETE("", playSongControl.DeletePlaylistSong)

	appColab := appPlaylist.Group("/collabs")
	appColab.Use(middleware.JWT([]byte(config.Cfg.TokenSecret)))
	appColab.POST("", collabControl.AddCollaboration)
	appColab.DELETE("", collabControl.DeleteCollaboration)

	appAlbumLike := appAlbum.Group("/like")
	appAlbumLike.Use(middleware.JWT([]byte(config.Cfg.TokenSecret)))
	appAlbumLike.POST("", albumLikeControl.AddAlbumLike)
	appAlbumLike.DELETE("", albumLikeControl.DeleteAlbumLike)

	appPlayActiv := appPlaylist.Group("/status")
	appPlayActiv.Use(middleware.JWT([]byte(config.Cfg.TokenSecret)))
	appPlayActiv.POST("", playActivControl.AddPlaylistActivity)
	appPlayActiv.DELETE("", playActivControl.DeletePlaylistActivity)
}
