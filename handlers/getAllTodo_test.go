package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
	"todo-list/mock"

	"github.com/gin-gonic/gin"
)

func TestGetAllTodo(t *testing.T) {
	server := gin.New()
	mockTodoService := &mock.MockTodoService{}
	mockTokenClient := mock.NewMockTokenClient()
	handler := NewHandler(mockTodoService, nil, &mockTokenClient)

	server.Handle(http.MethodGet, "/todos", handler.GetAllTodo)

	httpServer := httptest.NewServer(server)

	gin.SetMode(gin.TestMode)

	tests := map[string]struct {
		dbErr mock.ErrMock
		code  int
		want  gin.H
	}{
		"Successful retrieval": {
			dbErr: mock.OK,
			code:  http.StatusOK,
			want: gin.H{
				"result": []gin.H{{
					"todoId":    1,
					"title":     "Test Todo 1",
					"status":    "completed",
					"userId":    1,
					"createdAt": mock.MockTime.Format(time.RFC3339Nano),
					"updatedAt": mock.MockTime.Format(time.RFC3339Nano),
				},
				},
			},
		},
		"DB operation responded with error": {
			dbErr: mock.DBOperationError,
			code:  http.StatusInternalServerError,
			want:  gin.H{"message": "Internal server error"},
		},
	}

	for key, val := range tests {
		t.Run(key, func(t *testing.T) {
			mockTodoService.Err = val.dbErr
			client := http.Client{}
			reqURL := httpServer.URL + "/todos"
			req, err := http.NewRequest(http.MethodGet, reqURL, nil)

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
