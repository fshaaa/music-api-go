package albumRepository

import (
	"database/sql"
	"music-api-go/model"
)

type AlbumRepository interface {
	GetAllAlbums() ([]model.Albums, error)
	GetAlbum(id string) (model.Albums, error)
	AddAlbum(album model.Albums) error
	DeleteAlbum(id string) error
}

type albumRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) *albumRepository {
	return &albumRepository{db}
}

func (a *albumRepository) GetAllAlbums() ([]model.Albums, error) {
	var albums []model.Albums
	query := `SELECT * FROM albums LIMIT 10`

	row, err := a.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var album model.Albums
		err := row.Scan(&album.ID, &album.CreatedAt, &album.UpdatedAt, &album.Name, &album.Year, &album.Owner)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

func (a *albumRepository) GetAlbum(id string) (model.Albums, error) {
	var album model.Albums
	query := `SELECT * FROM albums WHERE id = $1`

	row, err := a.db.Query(query, id)
	if err != nil {
		return model.Albums{}, err
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&album.ID, &album.CreatedAt, &album.UpdatedAt, &album.Name, &album.Year, &album.Owner)
		if err != nil {
			return model.Albums{}, err
		}
	}
	return album, nil
}

func (a *albumRepository) AddAlbum(album model.Albums) error {
	query := `INSERT INTO albums VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := a.db.Exec(query, album.ID, album.CreatedAt, album.UpdatedAt, album.Name, album.Year, album.Owner)
	if err != nil {
		return err
	}
	return nil
}
func (a *albumRepository) DeleteAlbum(id string) error {
	query := `DELETE albums WHERE id = $1`
	_, err := a.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
