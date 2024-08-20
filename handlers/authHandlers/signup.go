package authHandlers

import (
	"net/http"
	"todo-list/models"
	"todo-list/services"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func Signup(c *gin.Context) {

	user := &models.User{}

	if err := c.ShouldBindJSON(user); err != nil {
		utils.HandleError(c, "Error while parsing json body", http.StatusUnprocessableEntity)
		return
	}

	//validate request body
	if utils.IsEmpty(user.Name) {
		utils.HandleError(c, "Username is required", http.StatusBadRequest)
		return
	}
	if utils.IsEmpty(user.Password) {
		utils.HandleError(c, "Password is required", http.StatusBadRequest)
		return
	}

	if !utils.IsLength(user.Name, 8, 20) {
		utils.HandleError(
			c,
			"Username must have atleast 8 and atmost 20 characters",
			http.StatusBadRequest,
		)
		return
	}
	if !utils.ValidPassword(user.Password) {
		utils.HandleError(
			c,
			"Password must have atleast one upper case and one digit and one special character",
			http.StatusBadRequest,
		)
		return
	}

	// insert the new user details
	err := services.AddUser(user)

	if err != nil {
		if err.(*pq.Error).Code == "23505" {
			utils.HandleError(c, "User already exists", http.StatusConflict)
			return
		}
		utils.HandleError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response{Message: "User signup successfull"}
	c.JSON(http.StatusOK, response)
}
