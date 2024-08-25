package todoHandlers

import (
	"errors"
	"net/http"
	"strconv"
	"todo-list/models"
	"todo-list/services"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId := c.GetInt("userId")

	storedTodo := &models.Todo{}
	err = services.GetTodo(todoId, userId, storedTodo) // retrieve todo from db

	if err != nil {
		if errors.Is(err, models.ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, storedTodo)
}
