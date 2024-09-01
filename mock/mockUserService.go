package mock

import (
	"database/sql"
	"todo-list/models"
	"todo-list/service"

	"golang.org/x/crypto/bcrypt"
)

type MockUserService struct {
	Err ErrMock
}

func (m *MockUserService) GetUser(user *models.User, storedUser *models.User) error {
	if m.Err == DBNotFoundInTodos {
		return service.ErrUserNotFound
	}

	if m.Err == DBOperationError {
		return service.ErrGetUser
	}

	storedUser.Name = "Kaushal"
	storedUser.Password = GenerateHashedPassword("K@ubb123b")
	storedUser.CreatedAt = MockTime
	storedUser.UpdatedAt = sql.NullTime{Valid: false, Time: MockTime}
	storedUser.UserId = 1

	return nil
}

func (m *MockUserService) AddUser(user *models.User) error {
	if m.Err == DBDuplicateEntry {
		return service.ErrDuplicateUniqueKey
	}

	if m.Err == DBOperationError {
		return service.ErrAddUser
	}

	return nil
}

func GenerateHashedPassword(password string) string {
	pw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pw)
}
