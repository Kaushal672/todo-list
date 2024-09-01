package handlers

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"todo-list/models"
	"todo-list/protogen/token"
	"todo-list/service"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handlers) Login(c *gin.Context) {

	user := &models.User{}

	if err := c.ShouldBindJSON(user); err != nil {
		if e, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, utils.FormatValidationError(e))
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json body"})
		}
		return
	}

	user.Name = strings.TrimSpace(user.Name)
	user.Password = strings.TrimSpace(user.Password)

	storedUser := &models.User{}
	// get the registered user data
	err := h.UserService.GetUser(user, storedUser) // handlers interface with struct,

	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Username or password is incorrect"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username or password is incorrect"})
		return
	}

	// token creation
	tokenString, err := h.TokengRPCClient.CreateToken(
		context.Background(),
		&token.UserId{UserId: int64(storedUser.UserId)},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	// create response body
	c.JSON(
		http.StatusOK,
		gin.H{"message": "User logged in successfully", "token": tokenString.GetToken()},
	)
}
