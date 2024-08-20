package todoHandlers

import (
	"encoding/json"
	"net/http"
	"todo-list/models"
	"todo-list/services"
	"todo-list/utils"
)

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.HandleError(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	userId := r.Context().Value("userId").(int)

	rows, err := services.GetAllTodo(userId)

	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result = []*models.Todo{}

	for rows.Next() {
		todo := new(models.Todo)
		rows.Scan(
			&todo.TodoId,
			&todo.Title,
			&todo.Status,
			&todo.UserId,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		result = append(result, todo)
	}

	jsonRes, err := json.Marshal(result)

	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}
