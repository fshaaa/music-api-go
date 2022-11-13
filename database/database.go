package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"music-api-go/config"
	"music-api-go/model"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

func InitDatabase() {
	cfg := config.Cfg

	addr := cfg.DB_ADDRESS
	port := cfg.DB_PORT
	user := cfg.DB_USERNAME
	pass := cfg.DB_PASSWORD
	name := cfg.DB_NAME

	psql := fmt.Sprintf("host=%s port=%s users=%s password=%s dbname=%s sslmode=disable",
		addr, port, user, pass, name)

	db, _ := gorm.Open(postgres.Open(psql), &gorm.Config{
		AllowGlobalUpdate: true,
	})

	DB = db

	DB.AutoMigrate(
		&model.Users{},
		&model.Playlists{},
		&model.Albums{},
		&model.AlbumLikes{},
		&model.Songs{},
		&model.PlaylistSongs{},
		&model.Collaborations{},
		&model.PlaylistActivities{},
	)
}

func InitDatabaseSql() *sql.DB {
	cfg := config.Cfg

	addr := cfg.DB_ADDRESS
	port := cfg.DB_PORT
	user := cfg.DB_USERNAME
	pass := cfg.DB_PASSWORD
	name := cfg.DB_NAME

	psql := fmt.Sprintf("host=%s port=%s users=%s password=%s dbname=%s sslmode=disable",
		addr, port, user, pass, name)

	db, err := sql.Open("postgres", psql)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
