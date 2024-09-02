package mock

import (
	"todo-list/models"
	"todo-list/service"

	"golang.org/x/crypto/bcrypt"
)

type MockUserService struct {
	Err ErrMock
}

func (m *MockUserService) GetUser(user *models.User) (*models.User, error) {
	if m.Err == DBNotFoundInUser {
		return nil, service.ErrUserNotFound
	}

	if m.Err == DBOperationError {
		return nil, service.ErrGetUser
	}

	return &models.User{
		Name:      "Kaushal",
		Password:  GenerateHashedPassword("K@ubb123b"),
		CreatedAt: &MockTime,
		UpdatedAt: &MockTime,
		UserId:    1,
	}, nil
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
