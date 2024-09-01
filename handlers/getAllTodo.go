package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetAllTodo(c *gin.Context) {

	userId := c.GetInt64("userId")

	result, err := h.TodoService.GetAllTodo(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
