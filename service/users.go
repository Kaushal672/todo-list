package service

import (
	"database/sql"
	"errors"
	"todo-list/models"

	"github.com/lib/pq"
)

var (
	ErrDuplicateUniqueKey = errors.New("duplicate unique key")
	ErrUserNotFound       = errors.New("user not found")
	ErrGetUser            = errors.New("unauthorized")
	ErrAddUser            = errors.New("could not register user")
)

type UserManager interface {
	GetUser(user *models.User) (*models.User, error)
	AddUser(user *models.User) error
}

type UserServices struct {
	DB *sql.DB
}

func (us *UserServices) GetUser(user *models.User) (*models.User, error) {
	query := "SELECT userId, username, password, createdAt, updatedAt FROM users WHERE username = $1"
	var resUser models.User
	err := us.DB.QueryRow(query, user.Name).
		Scan(&resUser.UserId,
			&resUser.Name,
			&resUser.Password,
			&resUser.CreatedAt,
			&resUser.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, ErrGetUser
	}

	return &resUser, nil
}

func (us *UserServices) AddUser(user *models.User) error {
	query := "INSERT INTO users(username, password) VALUES($1, $2)"
	_, err := us.DB.Exec(
		query,
		user.Name,
		user.Password,
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrDuplicateUniqueKey
		}
		return ErrAddUser
	}

	return nil
}
