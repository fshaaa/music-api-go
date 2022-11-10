package userRepository

import (
	"database/sql"
	"music-api-go/dto"
	"music-api-go/model"
	"music-api-go/util/setQuery"
)

type UserRepository interface {
	CreateUser(user model.Users) error
	LoginUser(user model.Users) (model.Users, error)
	GetUserById(id string) (model.Users, error)
	UpdateUser(id string, req map[string]interface{}) error
	DeleteUser(id string) error
	SearchUser(name string) ([]dto.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(user model.Users) error {
	query := `INSERT INTO users VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := u.db.Exec(query, user.ID, user.CreatedAt, user.UpdatedAt, user.Username, user.Password,
		user.Email, user.Fullname)
	return err
}

func (u *userRepository) LoginUser(user model.Users) (model.Users, error) {
	var userRes model.Users
	query := `SELECT * FROM users WHERE (username = $1 OR email = $2) AND password = $3`
	row, err := u.db.Query(query, user.Username, user.Email, user.Password)
	if err != nil {
		return userRes, err
	}
	defer row.Close()
	row.Next()
	err = row.Scan(&userRes.ID, &userRes.CreatedAt, &userRes.UpdatedAt, &userRes.Username, &userRes.Password,
		&userRes.Email, &userRes.Fullname)
	if err != nil {
		return userRes, err
	}
	return userRes, nil
}

func (u *userRepository) GetUserById(id string) (model.Users, error) {
	var userRes model.Users
	query := `SELECT * FROM users WHERE id = $1`
	row, err := u.db.Query(query, id)
	if err != nil {
		return userRes, err
	}
	defer row.Close()
	row.Next()
	err = row.Scan(&userRes.ID, &userRes.CreatedAt, &userRes.UpdatedAt, &userRes.Username, &userRes.Password,
		&userRes.Email, &userRes.Fullname)
	if err != nil {
		return userRes, err
	}
	return userRes, nil
}

func (u *userRepository) UpdateUser(id string, req map[string]interface{}) error {
	query, value := setQuery.UpdateDynamicQuery(req, "users", id)
	_, err := u.db.Exec(query, value...)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) SearchUser(name string) ([]dto.User, error) {
	var users []dto.User
	query := `SELECT id, username, email, fullname FROM users WHERE (username LIKE '%' || $1 || '%' OR
        email LIKE '%' || $2 || '%') LIMIT 10`
	row, err := u.db.Query(query, name, name)
	if err != nil {
		return nil, err
	}
	defer row.Close()
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
