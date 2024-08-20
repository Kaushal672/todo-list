package utils

import (
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, message string, status int) {
	c.JSON(status, gin.H{"message": message})
}
