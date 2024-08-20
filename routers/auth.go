package routers

import (
	"todo-list/handlers/authHandlers"
	"todo-list/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine) {
	authRouter := router.Group("/auth")
	authRouter.Use(middleware.AuthBodyValidator)
	authRouter.POST("/signup", authHandlers.Signup)
	authRouter.POST("/login", authHandlers.Login)
}
