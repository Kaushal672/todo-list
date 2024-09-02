package service

import (
	"database/sql"
	"errors"
	"log"
	"todo-list/models"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
	ErrGetTodo      = errors.New("could not get todo")
	ErrGetAllTodo   = errors.New("could not get all todos")
	ErrUpdateTodo   = errors.New("could not update todo")
	ErrDeleteTodo   = errors.New("could not delete todo")
	ErrAddTodo      = errors.New("could not add todo")
)

type TodoManager interface {
	AddTodo(todo *models.Todo, userId int64) error
	DeleteTodo(todoId int, userId int64) error
	GetTodo(todoId int, userId int64) (*models.Todo, error)
	UpdateTodo(todo *models.Todo, todoId int, userId int64) error
	GetAllTodo(userId int64) ([]models.Todo, error)
}

type TodoServices struct {
	DB *sql.DB
}

func (ts *TodoServices) AddTodo(todo *models.Todo, userId int64) error {
	query := "INSERT INTO todos(title, currentStatus, userId) VALUES($1, $2, $3)"
	_, err := ts.DB.Exec(
		query,
		todo.Title,
		todo.Status,
		userId,
	)

	if err != nil { // IF ONLY ERROR THEN RETURN ERROR
		return ErrAddTodo
	}

	return nil
}

func (ts *TodoServices) DeleteTodo(todoId int, userId int64) error {
	query := "DELETE FROM todos WHERE todoId = $1 AND userId = $2"
	_, err := ts.DB.Exec(
		query, // use local variables to store query
		todoId,
		userId,
	)

	if err != nil {
		return ErrDeleteTodo
	}

	return nil
}

func (ts *TodoServices) GetTodo(todoId int, userId int64) (*models.Todo, error) {
	query := "SELECT todoId, title, currentStatus, userId, createdAt, updatedAt FROM todos WHERE todoId = $1 AND userId = $2"
	var resTodo models.Todo
	err := ts.DB.QueryRow(query, todoId, userId).
		Scan(&resTodo.TodoId,
			&resTodo.Title,
			&resTodo.Status,
			&resTodo.UserId,
			&resTodo.CreatedAt,
			&resTodo.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrTodoNotFound
		}
		return nil, ErrGetTodo
	}

	return &resTodo, nil // check for err no rows, categorize the error, client should get generic error
}

func (ts *TodoServices) UpdateTodo(todo *models.Todo, todoId int, userId int64) error {
	query := "UPDATE todos SET title = $1, currentStatus = $2, updatedAt = CURRENT_TIMESTAMP WHERE todoId = $3 AND userId = $4"
	_, err := ts.DB.Exec(
		query,
		todo.Title,
		todo.Status,
		todoId,
		userId,
	)

	if err != nil {
		return ErrUpdateTodo
	}

	return nil
}

func (ts *TodoServices) GetAllTodo(userId int64) ([]models.Todo, error) {
	query := "SELECT todoId, title, currentStatus, userId, createdAt, updatedAt FROM todos WHERE userId = $1"
	rows, err := ts.DB.Query(
		query,
		userId,
	)
	if err != nil {
		return nil, ErrGetAllTodo
	}

	defer rows.Close()

	var result = []models.Todo{}

	for rows.Next() {
		todo := models.Todo{}
		err := rows.Scan(
			&todo.TodoId,
			&todo.Title,
			&todo.Status,
			&todo.UserId,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		if err != nil {
			log.Print(err.Error())
			return nil, ErrGetAllTodo
		}
		result = append(result, todo)
	}

	if rows.Err() != nil {
		return nil, ErrGetAllTodo
	}
	return result, nil
}
