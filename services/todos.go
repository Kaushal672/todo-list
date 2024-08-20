package services

import (
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

func UpdateTodo(todo *models.Todo, todoId int) error {
	_, err := database.DB.Exec(
		"UPDATE todos SET title = $1, currentStatus = $2, updatedAt = CURRENT_TIMESTAMP WHERE todoId = $3",
		todo.Title,
		todo.Status,
		todoId,
	)
	return err
}

func GetAllTodo(userId int) ([]*models.Todo, error) {
	rows, err := database.DB.Query(
		"SELECT todoId, title, currentStatus, userId, createdAt, updatedAt FROM todos WHERE userId = $1",
		userId,
	)

	if err != nil {
		return nil, err
	}

	var result = []*models.Todo{}

	for rows.Next() {
		todo := &models.Todo{}
		rows.Scan(
			&todo.TodoId,
			&todo.Title,
			&todo.Status,
			&todo.UserId,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		result = append(result, todo)
	}

	return result, nil
}
