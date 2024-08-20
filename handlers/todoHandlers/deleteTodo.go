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

func DeleteTodo(c *gin.Context) {

	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandleError(c, err.Error(), http.StatusBadRequest)
		return
	}

	storedTodo := &models.Todo{}
	err = services.GetTodo(todoId, storedTodo)
	if err == sql.ErrNoRows {
		utils.HandleError(c, "Todo not found", http.StatusNotFound)
		return
	} else if err != nil {
		utils.HandleError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	// get requested user's userId
	userId := c.GetInt("userId")

	// check to match the userId
	if userId != storedTodo.UserId {
		utils.HandleError(c, "Your not authorized to delete this todo", http.StatusForbidden)
		return
	}

	err = services.DeleteTodo(todoId)
	if err != nil {
		utils.HandleError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "Succesfully deleted todo"})
}
