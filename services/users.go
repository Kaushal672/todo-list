package services

import (
	"database/sql"
	"log"
	"todo-list/database"
	"todo-list/models"

	"github.com/lib/pq"
)

func GetUser(user *models.User, storedUser *models.User) error {
	err := database.DB.QueryRow("SELECT userId, username, password, createdAt, updatedAt FROM users WHERE username = $1", user.Name).
		Scan(&storedUser.UserId,
			&storedUser.Name,
			&storedUser.Password,
			&storedUser.CreatedAt,
			&storedUser.UpdatedAt)

	if err == nil {
		return nil
	}

	var e error
	if err == sql.ErrNoRows {
		e = models.ErrUserNotFound
	} else {
		e = models.ErrGetUser
	}

	log.Println(e)
	return e
}

func AddUser(user *models.User) error {
	_, err := database.DB.Exec(
		"INSERT INTO users(username, password) VALUES($1, $2)",
		user.Name,
		user.Password,
	)

	if err == nil {
		return nil
	}

	var e error
	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
		e = models.ErrDuplicateUniqueKey
	} else {
		e = models.ErrAddUser
	}

	log.Println(e)
	return e
}
