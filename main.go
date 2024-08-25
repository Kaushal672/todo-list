package main

import (
	"todo-list/handlers"
	"todo-list/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routers.SetupAuthRoutes(router)
	routers.SetupTodosRoutes(router)

	router.NoRoute(handlers.NotFound)

	router.Run(":3000")
}
