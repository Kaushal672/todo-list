package utils

import (
	"encoding/json"
	"net/http"
	"todo-list/models"
)

func HandleError(w http.ResponseWriter, message string, status int) {
	e := models.Error{Message: message}
	jsonRes, _ := json.Marshal(e)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonRes)
}
