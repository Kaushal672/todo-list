package main

import (
	"todo-list/database"
	"todo-list/handlers"
	"todo-list/routers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	database.ConnectToDB()
	router := gin.Default()

	routers.SetupAuthRoutes(router)
	routers.SetupTodosRoutes(router)

	router.NoRoute(handlers.NotFound)

	router.Run(":3000")
}
