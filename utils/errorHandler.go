package utils

import (
	"todo-list/models"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, message string, status int) {
	e := models.Error{Message: message}
	c.JSON(status, e)
}
