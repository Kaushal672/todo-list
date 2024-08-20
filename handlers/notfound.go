package handlers

import (
	"net/http"
	"todo-list/utils"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	utils.HandleError(w, "Not found", http.StatusNotFound)
}
