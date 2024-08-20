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

func UpdateTodo(c *gin.Context) {
	t, _ := c.Get("todo")
	todo := t.(*models.Todo)

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

	userId := c.GetInt("userId")

	//	check to match the userId
	if userId != storedTodo.UserId {
		utils.HandleError(c, "Your not authorized to update this todo", http.StatusForbidden)
		return
	}

	err = services.UpdateTodo(todo, todoId)

	if err != nil {
		utils.HandleError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "Todo updated successfully"})
}
