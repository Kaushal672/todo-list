package routers

import (
	"todo-list/handlers/authHandlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/signup", authHandlers.Signup)
		authRouter.POST("/login", authHandlers.Login)
	}
}
