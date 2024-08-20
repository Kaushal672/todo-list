package todoHandlers

import (
	"net/http"
	"todo-list/models"
	"todo-list/services"

	"todo-list/utils"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	t, _ := c.Get("todo")
	todo := t.(*models.Todo)

	userId := c.GetInt("userId")

	err := services.AddTodo(todo, userId)

	if err != nil {
		utils.HandleError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response{Message: "Todo created successfully"}
	c.JSON(http.StatusOK, response)
}
