package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IsAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	userTokenArr := strings.Split(tokenString, " ")

	if len(userTokenArr) < 2 {
		utils.HandleError(c, "Unauthorized", http.StatusUnauthorized)
		c.Abort()
		return
	}

	token, err := jwt.Parse(userTokenArr[1], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		utils.HandleError(c, "Unauthorized", http.StatusUnauthorized)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		utils.HandleError(c, "Unauthorized", http.StatusUnauthorized)
		c.Abort()
		return
	}

	userId := int(claims["userId"].(float64))
	c.Set("userId", userId)
	c.Next()
}
