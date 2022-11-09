package albumLikesRepository

import (
	"database/sql"
	"music-api-go/dto"
	"music-api-go/model"
)

type AlbumLikesRepository interface {
	AddAlbumLike(albumLike model.AlbumLikes) error
	DeleteAlbumLike(user_id, album_id string) error
	GetTotalAlbumLikes(album_id string) (int, error)
	GetUsersLikeAlbum(album_id string) ([]dto.User, error)
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

func (a *albumLikesRepository) GetTotalAlbumLikes(album_id string) (int, error) {
	var totalLikes = 0
	query := `SELECT COUNT(user_id) FROM album_likes WHERE album_id = $1`
	row, err := a.db.Query(query, album_id)
	if err != nil {
		return 0, err
	}
	if row.Next() {
		row.Scan(&totalLikes)
	}
	return totalLikes, nil
}

func (a *albumLikesRepository) GetUsersLikeAlbum(album_id string) ([]dto.User, error) {
	var users []dto.User
	query := `SELECT id, username, email, fullname FROM album_likes WHERE id = $1`
	row, err := a.db.Query(query, album_id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var user dto.User
		err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Fullname)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
