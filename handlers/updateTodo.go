package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"todo-list/models"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *Handlers) UpdateTodo(c *gin.Context) {
	todo := &models.Todo{}

	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid todo id"})
		return
	}

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

	err = h.TodoService.UpdateTodo(todo, todoId, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}
