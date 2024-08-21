package middleware

import (
	"net/http"
	"todo-list/models"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
)

func AuthBodyValidator(c *gin.Context) {
	user := &models.User{}

	if err := c.ShouldBindJSON(user); err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"message": "Error while parsing json body"},
		)
		return
	}

	if utils.IsEmpty(user.Name) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Username is required"})
		return
	}

	if utils.IsEmpty(user.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Password is required"})
		return
	}

	if !utils.IsLength(user.Name, 8, 20) {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": "Username must have atleast 8 and atmost 20 characters"},
		)
		return
	}

	if !utils.ValidPassword(user.Password) {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Password must be valid with atleast one upper case, one lower case, one special character(!, @, $, &), one digit with atleast 8 and atmost 16 characters",
			},
		)
		return
	}

	c.Set("user", user)
	c.Next()
}

func TodoBodyValidator(c *gin.Context) {
	todo := &models.Todo{}

	if err := c.ShouldBindJSON(todo); err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"message": "Error while parsing json body"},
		)
		return
	}

	//validate request body
	if utils.IsEmpty(todo.Title) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Title is required"})
		return
	}
	if utils.IsEmpty(todo.Status) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Status is required"})
		return
	}

	if !utils.IsLength(todo.Title, 5, 50) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Todo title too short"})
		return
	}
	if !utils.Contains([]string{"not_started", "in_progress", "completed"}, todo.Status) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid progress status"})
		return
	}

	c.Set("todo", todo)
	c.Next()
}
