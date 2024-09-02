package main

import (
	"log"
	"net/http"
	"todo-list/database"
	"todo-list/grpcClient"
	"todo-list/handlers"
	"todo-list/routers"
	"todo-list/service"
	"todo-list/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	// log.Println("init function invoked")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", validators.ValidatePassword)
	} // from binding get the gin validators
}

func main() {

	tokenClient, conn, err := grpcClient.CreateGrpcClient()

	if err != nil {
		log.Fatal("Error while creating gRPC client ", err)
	}

	defer conn.Close()

	router := gin.New()
	userService := &service.UserServices{DB: database.DB}
	todoService := &service.TodoServices{DB: database.DB}
	handler := handlers.NewHandler(todoService, userService, tokenClient)

	routers.SetupAuthRoutes(router, handler)
	routers.SetupTodosRoutes(router, handler)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	router.Run(":3000")
}
