package authHandlers

import (
	"database/sql"
	"net/http"
	"os"
	"time"
	"todo-list/models"
	"todo-list/services"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Login(c *gin.Context) {
	u, _ := c.Get("user")
	user := u.(*models.User)

	storedUser := &models.User{}
	// get the registered user data
	err := services.GetUser(user, storedUser)

	if err == sql.ErrNoRows { // check if user not found
		utils.HandleError(c, "Username or password is incorrect", http.StatusUnauthorized)
		return
	} else if err != nil {
		utils.HandleError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Password != storedUser.Password { // match the password
		utils.HandleError(c, "Username or password is incorrect", http.StatusUnauthorized)
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
		utils.HandleError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	// create response body
	response := models.AuthResponse{Message: "User logged in successfully", Token: tokenString}
	c.JSON(http.StatusOK, response)
}
