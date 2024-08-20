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

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		utils.HandleError(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	todo := new(models.Todo)
	err := json.NewDecoder(r.Body).Decode(&todo)
	defer r.Body.Close()
	if err != nil {
		utils.HandleError(w, "Error while parsing json body", http.StatusUnprocessableEntity)
		return
	}

	todo.TodoId, err = strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	//validate request body
	if utils.IsEmpty(todo.Title) {
		utils.HandleError(w, "Title is required", http.StatusBadRequest)
		return
	}
	if utils.IsEmpty(todo.Status) {
		utils.HandleError(w, "Status is required", http.StatusBadRequest)
		return
	}

	if !utils.IsLength(todo.Title, 5, 50) {
		utils.HandleError(w, "Todo title too short", http.StatusBadRequest)
		return
	}
	if !utils.Contains([]string{"not_started", "in_progress", "completed"}, todo.Status) {
		utils.HandleError(w, "Invalid progress status", http.StatusBadRequest)
		return
	}

	storedTodo := new(models.Todo)
	err = services.GetTodo(todo.TodoId, storedTodo)

	if err == sql.ErrNoRows {
		utils.HandleError(w, "Todo not found", http.StatusNotFound)
		return
	} else if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userId := r.Context().Value("userId").(int)

	//	check to match the userId
	if userId != storedTodo.UserId {
		utils.HandleError(w, "Your not authorized to update this todo", http.StatusForbidden)
		return
	}

	err = services.UpdateTodo(todo)

	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(models.Response{Message: "Todo updated successfully"})

	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
