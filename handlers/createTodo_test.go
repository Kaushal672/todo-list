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

	"github.com/gin-gonic/gin"
)

// tabular testing
func TestCreateTodo(t *testing.T) {
	server := gin.New()
	mockTodoService := &mock.MockTodoService{}
	mockTokenClient := mock.NewMockTokenClient()

	handler := NewHandler(mockTodoService, nil, &mockTokenClient)

	server.Handle(http.MethodPost, "/todos", handler.CreateTodo)

	httpServer := httptest.NewServer(server)

	gin.SetMode(gin.TestMode)

	tests := map[string]struct {
		dbErr       mock.ErrMock
		code        int
		requestBody gin.H
		want        gin.H
	}{
		"Successfull insertion of todo": {
			dbErr: mock.OK,
			code:  http.StatusOK,
			requestBody: gin.H{
				"title":  "Test Todo",
				"status": "completed",
			},
			want: gin.H{"message": "Todo created successfully"},
		},
		"Failed validation: Title field required": {
			dbErr:       mock.OK,
			code:        http.StatusBadRequest,
			requestBody: gin.H{"status": "completed"},
			want:        gin.H{"Title": "Title is required"},
		},
		"Failed validation: status field required": {
			dbErr:       mock.OK,
			code:        http.StatusBadRequest,
			requestBody: gin.H{"title": "Test todo"},
			want:        gin.H{"Status": "Status is required"},
		},
		"Failed validation: title min length": {
			dbErr: mock.OK,
			code:  http.StatusBadRequest,
			requestBody: gin.H{"title": "Test",
				"status": "completed"},
			want: gin.H{"Title": "Title must have atleast 5 characters"},
		},
		"Failed validation: title max length": {
			dbErr: mock.OK,
			code:  http.StatusBadRequest,
			requestBody: gin.H{
				"title":  "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum",
				"status": "completed",
			},
			want: gin.H{"Title": "Title must have atmost 50 characters"},
		},
		"Failed validation: status invalid value": {
			dbErr: mock.OK,
			code:  http.StatusBadRequest,
			requestBody: gin.H{"title": "Test todo",
				"status": "Completed"},
			want: gin.H{
				"Status": "Status should be one of the following: not_started in_progress completed",
			},
		},
		"Failed request due to invalid JSON": {
			dbErr: mock.OK,
			code:  http.StatusBadRequest,
			requestBody: gin.H{
				"title":  1234,
				"status": "completed",
			},
			want: gin.H{"message": "could not parse json body"},
		},
		"DB operation responded with error": {
			dbErr: mock.DBOperationError,
			code:  http.StatusInternalServerError,
			requestBody: gin.H{
				"title":  "Test Todo",
				"status": "completed",
			},
			want: gin.H{"message": "Internal server error"},
		},
	}

	for key, val := range tests { //TODO 3A testing
		t.Run(key, func(t *testing.T) {
			mockTodoService.Err = val.dbErr
			client := http.Client{}
			reqURL := httpServer.URL + "/todos"
			resJsonBody, err := json.Marshal(val.requestBody)
			if err != nil {
				t.Error("Error while marshalling response json body", err)
			}

			req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewReader(resJsonBody))
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
