package main

import (
	"log"
	"net/http"
	"todo-list/database"
	"todo-list/handlers"
	"todo-list/handlers/authHandlers"
	"todo-list/handlers/todoHandlers"
	"todo-list/middleware"

	_ "github.com/lib/pq"
)

func main() {
	database.ConnectToDB()
	// router := http.NewServeMux()

	http.HandleFunc("POST /auth/signup", authHandlers.Signup)
	http.HandleFunc("POST /auth/login", authHandlers.Login)

	// router.HandleFunc("POST /signup", authHandlers.Signup)
	// router.HandleFunc("POST /login", authHandlers.Login)
	// router.Handle("GET /todos", middleware.IsAuth(http.HandlerFunc(todoHandlers.GetAllTodo)))
	// router.Handle("POST /todos", middleware.IsAuth(http.HandlerFunc(todoHandlers.CreateTodo)))

	// router.Handle("GET /todos/{id}", middleware.IsAuth(http.HandlerFunc(todoHandlers.GetTodo)))
	// router.Handle("PUT /todos/{id}", middleware.IsAuth(http.HandlerFunc(todoHandlers.UpdateTodo)))
	// router.Handle("DELETE /todos/{id}", middleware.IsAuth(http.HandlerFunc(todoHandlers.DeleteTodo)))
	// authRouter := http.NewServeMux()
	// authRouter.Handle("/auth/", http.StripPrefix("/auth", router))
	//todosRouter := http.NewServeMux()
	// todosRouter.Handle("/todos/", http.StripPrefix("/todos", router))
	// server := http.Server{Addr: ":3000", Handler: router}
	//server.ListenAndServe()

	http.Handle("GET /todos", middleware.IsAuth(http.HandlerFunc(todoHandlers.GetAllTodo)))
	http.Handle("POST /todos", middleware.IsAuth(http.HandlerFunc(todoHandlers.CreateTodo)))

	http.Handle("GET /todos/{id}", middleware.IsAuth(http.HandlerFunc(todoHandlers.GetTodo)))
	http.Handle("PUT /todos/{id}", middleware.IsAuth(http.HandlerFunc(todoHandlers.UpdateTodo)))
	http.Handle("DELETE /todos/{id}", middleware.IsAuth(http.HandlerFunc(todoHandlers.DeleteTodo)))

	http.HandleFunc("/", handlers.NotFound)

	err := http.ListenAndServe("localhost:3000", nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}
