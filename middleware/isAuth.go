package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"todo-list/utils"

	"github.com/golang-jwt/jwt"
)

func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		userTokenArr := strings.Split(tokenString, " ")
		if len(userTokenArr) < 2 {
			utils.HandleError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(userTokenArr[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			utils.HandleError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			utils.HandleError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userId := int(claims["userId"].(float64))

		ctx := context.WithValue(r.Context(), "userId", userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
