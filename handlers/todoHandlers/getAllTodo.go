package todoHandlers

import (
	"net/http"
	"todo-list/services"

	"github.com/gin-gonic/gin"
)

func GetAllTodo(c *gin.Context) {

	userId := c.GetInt("userId")

	result, err := services.GetAllTodo(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
