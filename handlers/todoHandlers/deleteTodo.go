package todoHandlers

import (
	"net/http"
	"strconv"
	"todo-list/services"

	"github.com/gin-gonic/gin"
)

func DeleteTodo(c *gin.Context) {

	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get requested user's userId
	userId := c.GetInt("userId")

	err = services.DeleteTodo(todoId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succesfully deleted todo"})
}
