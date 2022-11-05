package albumLikesRepository

import (
	"database/sql"
	"music-api-go/model"
)

type AlbumLikesRepository interface {
	AddAlbumLike(albumLike model.AlbumLikes) error
	DeleteAlbumLike(user_id, album_id string) error
}

type albumLikesRepository struct {
	db *sql.DB
}

func NewAlbumLikesRepository(db *sql.DB) *albumLikesRepository {
	return &albumLikesRepository{db}
}

func (a *albumLikesRepository) AddAlbumLike(albumLike model.AlbumLikes) error {
	query := `INSERT INTO album_likes VALUES ($1,$2,$3,$4,$5)`
	_, err := a.db.Exec(query, albumLike.ID, albumLike.CreatedAt, albumLike.UpdatedAt, albumLike.User_id,
		albumLike.Album_id)
	if err != nil {
		return err
	}
	return nil
}

func (a *albumLikesRepository) DeleteAlbumLike(user_id, album_id string) error {
	query := `DELETE FROM album_likes WHERE user_id = $1 AND album_id = $2`
	_, err := a.db.Exec(query, user_id, album_id)
	if err != nil {
		return err
	}
	return nil
}
