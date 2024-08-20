package authHandlers

import (
	"encoding/json"
	"net/http"
	"todo-list/models"
	"todo-list/services"
	"todo-list/utils"

	"github.com/lib/pq"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.HandleError(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	user := new(models.User)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.HandleError(w, "Error while parsing json body", http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()

	//validate request body
	if utils.IsEmpty(user.Name) {
		utils.HandleError(w, "Username is required", http.StatusBadRequest)
		return
	}
	if utils.IsEmpty(user.Password) {
		utils.HandleError(w, "Password is required", http.StatusBadRequest)
		return
	}

	if !utils.IsLength(user.Name, 8, 20) {
		utils.HandleError(
			w,
			"Username must have atleast 8 and atmost 20 characters",
			http.StatusBadRequest,
		)
		return
	}
	if !utils.ValidPassword(user.Password) {
		utils.HandleError(
			w,
			"Password must have atleast one upper case and one digit and one special character",
			http.StatusBadRequest,
		)
		return
	}

	// insert the new user details
	err := services.AddUser(user)

	if err != nil {
		if err.(*pq.Error).Code == "23505" {
			utils.HandleError(w, "User already exists", http.StatusConflict)
			return
		}
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response{Message: "User signup successfull"}
	jsonRes, err := json.Marshal(response)
	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}
