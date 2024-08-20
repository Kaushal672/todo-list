package todoHandlers

import (
	"net/http"
	"todo-list/models"
	"todo-list/services"

	"todo-list/utils"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {

	todo := &models.Todo{}

	if err := c.ShouldBindJSON(todo); err != nil {
		utils.HandleError(c, "Error while parsing json body", http.StatusUnprocessableEntity)
		return
	}

	//validate request body
	if utils.IsEmpty(todo.Title) {
		utils.HandleError(c, "Title is required", http.StatusBadRequest)
		return
	}
	if utils.IsEmpty(todo.Status) {
		utils.HandleError(c, "Status is required", http.StatusBadRequest)
		return
	}

	if !utils.IsLength(todo.Title, 5, 50) {
		utils.HandleError(c, "Todo title too short", http.StatusBadRequest)
		return
	}
	if !utils.Contains([]string{"not_started", "in_progress", "completed"}, todo.Status) {
		utils.HandleError(c, "Invalid progress status", http.StatusBadRequest)
		return
	}

	userId := c.GetInt("userId")

	err := services.AddTodo(todo, userId)

	if err != nil {
		utils.HandleError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response{Message: "Todo created successfully"}
	c.JSON(http.StatusOK, response)
}
