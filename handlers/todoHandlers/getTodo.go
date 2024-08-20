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

func GetTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.HandleError(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	// get todoId from route param
	todoId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	storedTodo := new(models.Todo)
	err = services.GetTodo(todoId, storedTodo) // retrieve todo from db

	if err == sql.ErrNoRows {
		utils.HandleError(w, "Todo not found", http.StatusNotFound)
		return
	} else if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userId := r.Context().Value("userId").(int)

	if userId != storedTodo.UserId {
		utils.HandleError(w, "Your not authorized to view this todo", http.StatusForbidden)
		return
	}

	jsonRes, err := json.Marshal(storedTodo)

	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}
