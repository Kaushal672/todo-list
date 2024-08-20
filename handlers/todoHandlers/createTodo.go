package todoHandlers

import (
	"encoding/json"
	"net/http"
	"todo-list/models"
	"todo-list/services"

	"todo-list/utils"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
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

	userId := r.Context().Value("userId").(int)

	err = services.AddTodo(todo, userId)

	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(models.Response{Message: "Todo created successfully"})

	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
