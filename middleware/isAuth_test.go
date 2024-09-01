package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"todo-list/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func TestIsAuth(t *testing.T) {
	secret := "my_secret_key"
	os.Setenv("JWT_SECRET", secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": 1,
	})
	tokenString, _ := token.SignedString([]byte(secret))

	server := gin.New()
	mockTokenClient := mock.NewMockTokenClient()

	server.Use(IsAuth(&mockTokenClient))

	server.Handle(http.MethodGet, "/test", func(c *gin.Context) {
		userId, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"userId": userId})
	})

	httpServer := httptest.NewServer(server)

	gin.SetMode(gin.TestMode)

	tests := map[string]struct {
		tokenErr            mock.ErrMock
		authorizationHeader string
		code                int
		want                gin.H
	}{
		"No Authorization Header": {
			tokenErr:            mock.StatusOkInTokenService,
			authorizationHeader: "",
			code:                http.StatusUnauthorized,
			want:                gin.H{"message": "Unauthorized"},
		},
		"Malformed Authorization Header": {
			tokenErr:            mock.StatusOkInTokenService,
			authorizationHeader: "Bearer",
			code:                http.StatusUnauthorized,
			want:                gin.H{"message": "Unauthorized"},
		},
		"Invalid Token": {
			tokenErr:            mock.ErrInTokenVerification,
			authorizationHeader: "Bearer invalidtoken",
			code:                http.StatusUnauthorized,
			want:                gin.H{"message": "Unauthorized"},
		},
		"Valid Token": {
			tokenErr:            mock.StatusOkInTokenService,
			authorizationHeader: "Bearer " + tokenString,
			code:                http.StatusOK,
			want:                gin.H{"userId": 1},
		},
	}
	for key, val := range tests {
		t.Run(key, func(t *testing.T) {
			mockTokenClient.Err = val.tokenErr
			client := http.Client{}
			reqURL := httpServer.URL + "/test"

			req, _ := http.NewRequest(http.MethodGet, reqURL, nil)
			if val.authorizationHeader != "" {
				req.Header.Set("Authorization", val.authorizationHeader)
			}
			res, err := client.Do(req)

			if err != nil {
				t.Error("Error while sending request", err)
			}

			body, err := io.ReadAll(res.Body)

			if err != nil {
				t.Error("Error while reading body", err)
			}

			var resBody gin.H
			err = json.Unmarshal(body, &resBody)

			if err != nil {
				t.Error("Error while unmarshalling response body", err)
			}

			if res.StatusCode != val.code {
				t.Errorf("Expected status code %d, got %d", val.code, res.StatusCode)
			}

			if !reflect.DeepEqual(resBody, val.want) {
				if fmt.Sprint(resBody) != fmt.Sprint(val.want) {
					t.Errorf(
						"Expected response body to contain %s, got %s",
						val.want,
						resBody,
					)
				}
			}
		})
	}
}
