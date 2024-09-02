package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"todo-list/service"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid todo id"})
		return
	}

	userId := c.GetInt64("userId")

	todo, err := h.TodoService.GetTodo(todoId, userId)

	if err != nil {
		if errors.Is(err, service.ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": todo})
}
