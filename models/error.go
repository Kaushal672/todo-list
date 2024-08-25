package models

import "errors"

var (
	ErrDuplicateUniqueKey = errors.New("duplicate unique key")
	ErrUserNotFound       = errors.New("user not found")
	ErrTodoNotFound       = errors.New("todo not found")
	ErrGetUser            = errors.New("unauthorized")
	ErrAddUser            = errors.New("could not register user")
	ErrGetTodo            = errors.New("could not get todo")
	ErrGetAllTodo         = errors.New("could not get all todos")
	ErrUpdateTodo         = errors.New("could not update todo")
	ErrDeleteTodo         = errors.New("could not delete todo")
	ErrAddTodo            = errors.New("could not delete todo")
)
