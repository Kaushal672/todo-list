package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"todo-list/mock"
	"todo-list/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", validators.ValidatePassword)
	} // from binding get the gin validators
}

func TestLogin(t *testing.T) {
	server := gin.New()
	mockUserService := &mock.MockUserService{}
	mockTokenClient := mock.NewMockTokenClient()
	handler := NewHandler(nil, mockUserService, &mockTokenClient)

	server.Handle(http.MethodPost, "/auth/login", handler.Login)

	httpServer := httptest.NewServer(server)

	gin.SetMode(gin.TestMode)

	tests := map[string]struct {
		requestBody gin.H
		dbErr       mock.ErrMock
		tokenErr    mock.ErrMock
		code        int
		want        gin.H
	}{
		"Successful login": {
			requestBody: gin.H{
				"name":     "Kaushal",
				"password": "K@ubb123b",
			},
			dbErr:    mock.OK,
			tokenErr: mock.StatusOkInTokenService,
			code:     http.StatusOK,
			want:     gin.H{"message": "User logged in successfully"},
		},
		"User not found": {
			requestBody: gin.H{
				"name":     "unknown",
				"password": "K@ubb123",
			},
			dbErr:    mock.DBNotFoundInUser,
			tokenErr: mock.StatusOkInTokenService,
			code:     http.StatusUnauthorized,
			want:     gin.H{"message": "Username or password is incorrect"},
		},
		"Failed validation: name is required": {
			requestBody: gin.H{"password": "K@ubb123b"},
			dbErr:       mock.OK,
			tokenErr:    mock.StatusOkInTokenService,
			code:        http.StatusBadRequest,
			want:        gin.H{"Name": "Name is required"},
		},
		"Failed validation: password is required": {
			requestBody: gin.H{"name": "Kaushal"},
			dbErr:       mock.OK,
			tokenErr:    mock.StatusOkInTokenService,
			code:        http.StatusBadRequest,
			want:        gin.H{"Password": "Password is required"},
		},
		"Failed validation: name min length": {
			requestBody: gin.H{"name": "test", "password": "K@ubb123b"},
			dbErr:       mock.OK,
			tokenErr:    mock.StatusOkInTokenService,
			code:        http.StatusBadRequest,
			want:        gin.H{"Name": "Name must have atleast 5 characters"},
		},
		"Failed validation: name max length": {
			requestBody: gin.H{
				"name":     "testkajldhfkjadbfkjahdkfjakdfhjkdbj",
				"password": "K@ubb123b",
			},
			dbErr:    mock.OK,
			tokenErr: mock.StatusOkInTokenService,
			code:     http.StatusBadRequest,
			want:     gin.H{"Name": "Name must have atmost 16 characters"},
		},
		"Failed validation: password min length": {
			requestBody: gin.H{"name": "test user", "password": "K@ub1"},
			dbErr:       mock.OK,
			tokenErr:    mock.StatusOkInTokenService,
			code:        http.StatusBadRequest,
			want: gin.H{
				"Password": "Password must include characters between 8 and 16, one uppercase, one lowercase, one digit and one special character",
			},
		},
		"Failed validation: password max length": {
			requestBody: gin.H{
				"name":     "test user",
				"password": "K@ub1232jhdjfbajhbfjjdkfbakjfbakjnsj",
			},
			dbErr:    mock.OK,
			tokenErr: mock.StatusOkInTokenService,
			code:     http.StatusBadRequest,
			want: gin.H{
				"Password": "Password must include characters between 8 and 16, one uppercase, one lowercase, one digit and one special character",
			},
		},
		"Failed validation: password missing upper case": {
			requestBody: gin.H{
				"name":     "test user",
				"password": "k@ub1232h",
			},
			dbErr:    mock.OK,
			tokenErr: mock.StatusOkInTokenService,
			code:     http.StatusBadRequest,
			want: gin.H{
				"Password": "Password must include characters between 8 and 16, one uppercase, one lowercase, one digit and one special character",
			},
		},
		"Failed validation: password missing lower case": {
			requestBody: gin.H{
				"name":     "test user",
				"password": "K@UB1232",
			},
			dbErr:    mock.OK,
			tokenErr: mock.StatusOkInTokenService,
			code:     http.StatusBadRequest,
			want: gin.H{
				"Password": "Password must include characters between 8 and 16, one uppercase, one lowercase, one digit and one special character",
			},
		},
		"Failed validation: password missing digit": {
			requestBody: gin.H{
				"name":     "test user",
				"password": "K@UBSsajkd",
			},
			dbErr:    mock.OK,
			tokenErr: mock.StatusOkInTokenService,
			code:     http.StatusBadRequest,
			want: gin.H{
				"Password": "Password must include characters between 8 and 16, one uppercase, one lowercase, one digit and one special character",
			},
		},
		"Failed validation: password missing special character": {
			requestBody: gin.H{
				"name":     "test user",
				"password": "K1UBSsajkd",
			},
			dbErr:    mock.OK,
			tokenErr: mock.StatusOkInTokenService,
			code:     http.StatusBadRequest,
			want: gin.H{
				"Password": "Password must include characters between 8 and 16, one uppercase, one lowercase, one digit and one special character",
			},
		},
		"Failed token creation": {
			requestBody: gin.H{
				"name":     "test user",
				"password": "K@ubb123b",
			},
			dbErr:    mock.OK,
			tokenErr: mock.ErrInTokenCreation,
			code:     http.StatusInternalServerError,
			want: gin.H{
				"message": "Internal server error",
			},
		},
		"Incorrect Password": {
			requestBody: gin.H{
				"name":     "Kaushal",
				"password": "K@ubb123",
			},
			dbErr:    mock.OK,
			tokenErr: mock.StatusOkInTokenService,
			code:     http.StatusUnauthorized,
			want:     gin.H{"message": "Username or password is incorrect"},
		},
		"Failed user retrieval": {
			requestBody: gin.H{
				"name":     "Kaushal",
				"password": "K@ubb123",
			},
			dbErr:    mock.DBOperationError,
			tokenErr: mock.StatusOkInTokenService,
			code:     http.StatusInternalServerError,
			want:     gin.H{"message": "Internal server error"},
		},
	}

	for key, val := range tests {
		t.Run(key, func(t *testing.T) {
			mockTokenClient.Err = val.tokenErr
			mockUserService.Err = val.dbErr
			client := http.Client{}
			reqURL := httpServer.URL + "/auth/login"
			resJsonBody, err := json.Marshal(val.requestBody)
			if err != nil {
				t.Error("Error while marshalling response json body", err)
			}
			req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(resJsonBody))

			if err != nil {
				t.Error("Error while creating request", err)
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

			if res.StatusCode == http.StatusOK {
				if _, ok := resBody["token"]; !ok {
					t.Errorf("Expected a token in the response, but none was found")
				}
			}

			if !reflect.DeepEqual(resBody["message"], val.want["message"]) {
				if fmt.Sprint(resBody["message"]) != fmt.Sprint(val.want["message"]) {
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
