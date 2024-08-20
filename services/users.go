package services

import (
	"todo-list/database"
	"todo-list/models"
)

func GetUser(user *models.User, storedUser *models.User) error {
	err := database.DB.QueryRow("SELECT userId, username, password, createdAt, updatedAt FROM users WHERE username = $1", user.Name).
		Scan(&storedUser.UserId,
			&storedUser.Name,
			&storedUser.Password,
			&storedUser.CreatedAt,
			&storedUser.UpdatedAt)
	return err
}

func AddUser(user *models.User) error {
	_, err := database.DB.Exec(
		"INSERT INTO users(username, password) VALUES($1, $2)",
		user.Name,
		user.Password,
	)
	return err
}
