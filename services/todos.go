package services

import (
	"database/sql"
	"log"
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

	if err == nil {
		return nil
	}

	return models.ErrAddTodo
}

func DeleteTodo(todoId int, userId int) error {
	_, err := database.DB.Exec(
		"DELETE FROM todos WHERE todoId = $1 AND userId = $2",
		todoId,
		userId,
	)

	if err == nil {
		return nil
	}

	return models.ErrDeleteTodo
}

func GetTodo(todoId int, userId int, storedTodo *models.Todo) error {
	err := database.DB.QueryRow("SELECT todoId, title, currentStatus, userId, createdAt, updatedAt FROM todos WHERE todoId = $1 AND userId = $2", todoId, userId).
		Scan(&storedTodo.TodoId,
			&storedTodo.Title,
			&storedTodo.Status,
			&storedTodo.UserId,
			&storedTodo.CreatedAt,
			&storedTodo.UpdatedAt)

	if err == nil {
		return nil
	}

	var e error
	if err == sql.ErrNoRows {
		e = models.ErrTodoNotFound
	} else {
		e = models.ErrGetTodo
	}
	log.Println(e)
	return e // check for err no rows, categorize the error, client should get generic error
}

func UpdateTodo(todo *models.Todo, todoId int, userId int) error {
	_, err := database.DB.Exec(
		"UPDATE todos SET title = $1, currentStatus = $2, updatedAt = CURRENT_TIMESTAMP WHERE todoId = $3 AND userId = $4",
		todo.Title,
		todo.Status,
		todoId,
		userId,
	)

	if err == nil {
		return nil
	}

	return models.ErrUpdateTodo
}

func GetAllTodo(userId int) ([]*models.Todo, error) {
	rows, err := database.DB.Query(
		"SELECT todoId, title, currentStatus, userId, createdAt, updatedAt FROM todos WHERE userId = $1",
		userId,
	)

	if err != nil {
		return nil, models.ErrGetAllTodo
	}

	defer rows.Close()

	var result = []*models.Todo{}

	for rows.Next() {
		todo := &models.Todo{}
		err := rows.Scan(
			&todo.TodoId,
			&todo.Title,
			&todo.Status,
			&todo.UserId,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		if err != nil {
			return nil, models.ErrGetAllTodo
		}
		result = append(result, todo)
	}

	if rows.Err() != nil {
		return nil, models.ErrGetAllTodo
	}
	return result, nil
}
