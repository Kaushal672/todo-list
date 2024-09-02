package middleware

import (
	"context"
	"net/http"
	"strings"
	"token-management-service/protogen/token"

	"github.com/gin-gonic/gin"
)

func IsAuth(tokengRPCClient token.TokenClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		userTokenArr := strings.Split(tokenString, " ")

		if len(userTokenArr) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		res, err := tokengRPCClient.VerifyToken(
			context.Background(),
			&token.TokenString{Token: userTokenArr[1]},
		)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		c.Set("userId", res.GetUserId())

		c.Next()
	}
}
