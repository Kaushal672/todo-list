package todoHandlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"todo-list/models"
	"todo-list/services"
	"todo-list/utils"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		utils.HandleError(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	todoId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	storedTodo := new(models.Todo)
	err = services.GetTodo(todoId, storedTodo)
	if err == sql.ErrNoRows {
		utils.HandleError(w, "Todo not found", http.StatusNotFound)
		return
	} else if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get requested user's userId
	userId := r.Context().Value("userId").(int)

	// check to match the userId
	if userId != storedTodo.UserId {
		utils.HandleError(w, "Your not authorized to delete this todo", http.StatusForbidden)
		return
	}

	err = services.DeleteTodo(todoId)
	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(models.Response{Message: "Succesfully deleted todo"})
	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
