package mock

import (
	"time"
	"todo-list/models"
	"todo-list/service"
)

// TODO create new package for mock
type ErrMock int

const (
	DBOperationError ErrMock = iota
	OK
	DBNotFoundInTodos
	DBNotFoundInUser
	DBDuplicateEntry
)

var (
	MockTime = time.Date(2024, 8, 28, 10, 45, 53, int(time.Second), time.UTC)
)

type MockTodoService struct {
	Err ErrMock
}

func (s *MockTodoService) AddTodo(todo *models.Todo, userId int64) error {
	if s.Err == DBOperationError {
		return service.ErrAddTodo
	}
	return nil
}

func (s *MockTodoService) DeleteTodo(todoId int, userId int64) error {
	if s.Err == DBOperationError {
		return service.ErrDeleteTodo
	}
	return nil
}

func (s *MockTodoService) GetTodo(todoId int, userId int64, storedTodo *models.Todo) error {
	if s.Err == DBNotFoundInTodos {
		return service.ErrTodoNotFound
	}
	if s.Err == DBOperationError {
		return service.ErrGetTodo
	}

	storedTodo.Title = "Test Todo 1"
	storedTodo.Status = "completed"
	storedTodo.UserId = 1
	storedTodo.TodoId = 1
	storedTodo.CreatedAt = MockTime
	storedTodo.UpdatedAt = &MockTime

	return nil
}

func (s *MockTodoService) UpdateTodo(todo *models.Todo, todoId int, userId int64) error {
	if s.Err == DBOperationError {
		return service.ErrUpdateTodo
	}
	return nil
}

func (s *MockTodoService) GetAllTodo(userId int64) ([]models.Todo, error) {
	if s.Err == DBOperationError {
		return nil, service.ErrGetAllTodo
	}

	return []models.Todo{
		{
			TodoId:    1,
			Title:     "Test Todo 1",
			Status:    "completed",
			UserId:    1,
			CreatedAt: MockTime,
			UpdatedAt: &MockTime,
		},
	}, nil
}
