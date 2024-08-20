package services

import (
	"database/sql"
	"todo-list/database"
	"todo-list/models"
)

func AddTodo(todo *models.Todo, userId int) error {
	_, err := database.DB.Exec(
		"INSERT INTO todos(title, currentStatus, userId) VALUES($1, $2, $3)",
		todo.Title,
		todo.Status,
		userId,
	)
	return err
}

func DeleteTodo(todoId int) error {
	_, err := database.DB.Exec("DELETE FROM todos WHERE todoId = $1", todoId)
	return err
}

func GetTodo(todoId int, storedTodo *models.Todo) error {
	err := database.DB.QueryRow("SELECT todoId, title, currentStatus, userId, createdAt, updatedAt FROM todos WHERE todoId = $1", todoId).
		Scan(&storedTodo.TodoId,
			&storedTodo.Title,
			&storedTodo.Status,
			&storedTodo.UserId,
			&storedTodo.CreatedAt,
			&storedTodo.UpdatedAt)
	return err
}

func UpdateTodo(todo *models.Todo) error {
	_, err := database.DB.Exec(
		"UPDATE todos SET title = $1, currentStatus = $2, updatedAt = CURRENT_TIMESTAMP WHERE todoId = $3",
		todo.Title,
		todo.Status,
		todo.TodoId,
	)
	return err
}

func GetAllTodo(userId int) (*sql.Rows, error) {
	rows, err := database.DB.Query(
		"SELECT todoId, title, currentStatus, userId, createdAt, updatedAt FROM todos WHERE userId = $1",
		userId,
	)
	return rows, err
}
