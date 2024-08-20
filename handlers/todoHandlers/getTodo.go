package todoHandlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"todo-list/models"
	"todo-list/services"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	// get todoId from route param
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandleError(c, err.Error(), http.StatusBadRequest)
		return
	}

	storedTodo := &models.Todo{}
	err = services.GetTodo(todoId, storedTodo) // retrieve todo from db

	if err == sql.ErrNoRows {
		utils.HandleError(c, "Todo not found", http.StatusNotFound)
		return
	} else if err != nil {
		utils.HandleError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	userId := c.GetInt("userId")

	if userId != storedTodo.UserId {
		utils.HandleError(c, "Your not authorized to view this todo", http.StatusForbidden)
		return
	}

	c.JSON(http.StatusOK, storedTodo)
}
