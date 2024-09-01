package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) DeleteTodo(c *gin.Context) {

	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid todo id"})
		return
	}

	userId := c.GetInt64("userId")

	err = h.TodoService.DeleteTodo(todoId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succesfully deleted todo"})
}
