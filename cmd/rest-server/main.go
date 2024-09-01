package main

import (
	"log"
	"net/http"
	"todo-list/database"
	"todo-list/handlers"
	"todo-list/protogen/token"
	"todo-list/routers"
	"todo-list/service"
	"todo-list/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	// log.Println("init function invoked")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", validators.ValidatePassword)
	} // from binding get the gin validators
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Error in client connection", err)
	}
	defer conn.Close()

	t := token.NewTokenClient(conn)

	router := gin.New()
	userService := &service.UserServices{DB: database.DB}
	todoService := &service.TodoServices{DB: database.DB}
	handler := handlers.NewHandler(todoService, userService, t)

	routers.SetupAuthRoutes(router, handler)
	routers.SetupTodosRoutes(router, handler)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	router.Run(":3000")
}
