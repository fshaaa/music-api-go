package main

import (
	"github.com/labstack/echo/v4"
	"music-api-go/config"
	"music-api-go/database"
	"music-api-go/route"
)

func main() {
	config.InitConfig()
	database.InitDatabase()

	db := database.InitDatabaseSql()
	app := echo.New()
	route.NewRoute(app, db)
	defer db.Close()

	app.Logger.Fatal(app.Start(":8080"))
}
