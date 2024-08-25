package authHandlers

import (
	"errors"
	"net/http"
	"strings"
	"todo-list/models"
	"todo-list/services"
	"todo-list/utils"
	"todo-list/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
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

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	// replace raw password with hashed password
	user.Password = string(hash)

	// Insert the new user details
	err = services.AddUser(user)

	if err != nil {
		if errors.Is(err, models.ErrDuplicateUniqueKey) {
			c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User signup successful"})
}
