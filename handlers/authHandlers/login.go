package authHandlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"time"
	"todo-list/models"
	"todo-list/services"
	"todo-list/utils"

	"github.com/golang-jwt/jwt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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
			"Password must be valid with atleast one upper case, one lower case, one special character(!, @, $, &), one digit with atleast 8 and atmost 16 characters",
			http.StatusBadRequest,
		)
		return
	}

	storedUser := new(models.User)
	// get the registered user data
	err := services.GetUser(user, storedUser)

	if err == sql.ErrNoRows { // check if user not found
		utils.HandleError(w, "Username or password is incorrect", http.StatusNotFound)
		return
	} else if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Password != storedUser.Password { // match the password
		utils.HandleError(w, "Username or password is incorrect", http.StatusUnauthorized)
		return
	}

	// token creation
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": storedUser.UserId,
		"nbf":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Minute * 10).Unix(),
		"iat":    time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create response body
	response := models.AuthResponse{Message: "User logged in successfully", Token: tokenString}
	jsonRes, err := json.Marshal(response)

	if err != nil {
		utils.HandleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonRes)
}
