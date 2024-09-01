package routers

import (
	"todo-list/handlers"
	"todo-list/middleware"

	"github.com/gin-gonic/gin"
)

func SetupTodosRoutes(router *gin.Engine, handler *handlers.Handlers) {

	todoRouter := router.Group("/todos")
	todoRouter.Use(middleware.IsAuth(handler.TokengRPCClient))
	{
		todoRouter.GET("/", handler.GetAllTodo)
		todoRouter.POST("/", handler.CreateTodo)
		todoRouter.GET("/:id", handler.GetTodo)
		todoRouter.PUT("/:id", handler.UpdateTodo)
		todoRouter.DELETE("/:id", handler.DeleteTodo)
	}
}
