package routers

import (
	"todo-list/handlers/todoHandlers"
	"todo-list/middleware"

	"github.com/gin-gonic/gin"
)

func SetupTodosRoutes(router *gin.Engine) {
	todoRouter := router.Group("/todos")
	todoRouter.Use(middleware.IsAuth)
	{
		todoRouter.GET("/", todoHandlers.GetAllTodo)
		todoRouter.POST("/", todoHandlers.CreateTodo)
		todoRouter.GET("/:id", todoHandlers.GetTodo)
		todoRouter.PUT("/:id", todoHandlers.UpdateTodo)
		todoRouter.DELETE("/:id", todoHandlers.DeleteTodo)
	}
}
