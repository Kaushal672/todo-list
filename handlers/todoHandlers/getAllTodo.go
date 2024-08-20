package todoHandlers

import (
	"net/http"
	"todo-list/services"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
)

func GetAllTodo(c *gin.Context) {

	userId := c.GetInt("userId")

	result, err := services.GetAllTodo(userId)

	if err != nil {
		utils.HandleError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
