package handlers

import (
	"net/http"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	utils.HandleError(c, "Not found", http.StatusNotFound)
}
