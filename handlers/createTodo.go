package handlers

import (
	"net/http"
	"strings"
	"todo-list/models"
	"todo-list/protogen/token"
	"todo-list/service"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	TodoService     service.TodoManager
	UserService     service.UserManager
	TokengRPCClient token.TokenClient
}

func NewHandler(
	todoService service.TodoManager,
	userService service.UserManager,
	tokengRPCClient token.TokenClient,
) *Handlers {
	return &Handlers{
		TodoService:     todoService,
		UserService:     userService,
		TokengRPCClient: tokengRPCClient,
	}
}

func (h *Handlers) CreateTodo(c *gin.Context) {
	todo := &models.Todo{}

	if err := c.ShouldBindJSON(todo); err != nil {
		if e, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, utils.FormatValidationError(e))
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json body"})
		}
		return
	}

	todo.Title = strings.TrimSpace(todo.Title)
	todo.Status = strings.TrimSpace(todo.Status)

	userId := c.GetInt64("userId")

	err := h.TodoService.AddTodo(todo, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo created successfully"})
}
