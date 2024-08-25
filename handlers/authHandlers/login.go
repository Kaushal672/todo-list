package authHandlers

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"
	"todo-list/models"
	"todo-list/services"
	"todo-list/utils"
	"todo-list/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	user := &models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json body"})
		return
	}

	user.Name = strings.TrimSpace(user.Name)
	user.Password = strings.TrimSpace(user.Password)

	if err := validators.Validate.Struct(user); err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, utils.FormatValidationError(errs))
		return
	}

	storedUser := &models.User{}
	// get the registered user data
	err := services.GetUser(user, storedUser)

	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Username or password is incorrect"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username or password is incorrect"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// create response body
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": tokenString})
}
