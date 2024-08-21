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
	u, _ := c.Get("user")
	user := u.(*models.User)

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

	c.JSON(http.StatusOK, gin.H{"message": "User signup successfull"})
}
