package handlers

import (
	"errors"
	"net/http"
	"strings"
	"todo-list/models"
	"todo-list/service"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handlers) Signup(c *gin.Context) {
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

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	// replace raw password with hashed password
	user.Password = string(hash)

	// Insert the new user details
	err = h.UserService.AddUser(user)

	if err != nil {
		if errors.Is(err, service.ErrDuplicateUniqueKey) {
			c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User signup successfull"})
}
