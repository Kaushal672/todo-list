package todoHandlers

import (
	"net/http"
	"strings"
	"todo-list/models"
	"todo-list/services"
	"todo-list/utils"
	"todo-list/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateTodo(c *gin.Context) {
	todo := &models.Todo{}

	if err := c.ShouldBindJSON(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json body"})
		return
	}

	todo.Title = strings.TrimSpace(todo.Title)
	todo.Status = strings.TrimSpace(todo.Status)

	if err := validators.Validate.Struct(todo); err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, utils.FormatValidationError(errs))
		return
	}

	userId := c.GetInt("userId")

	err := services.AddTodo(todo, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo created successfully"})
}
