package routers

import (
	"todo-list/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, handler *handlers.Handlers) {
	// first service instance

	// handler instance with that service
	// handler := handlers.NewHandler(nil, userService) // in main

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/signup", handler.Signup)
		authRouter.POST("/login", handler.Login)
	}
}
