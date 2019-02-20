package main

import (
	"net/http"

	"github.com/NullJupiter/GoTodoApp/src/data_api/db"
	"github.com/NullJupiter/GoTodoApp/src/data_api/routing"
)

func main() {
	// Initialize database instance
	err := db.InitDB("postgres://localhost/todoapp?sslmode=disable")
	if err != nil {
		panic("Could not initialize a database instance!")
	}
	defer db.GetDB().Close()

	// Initialize routing and get router
	router := routing.InitRouting()

	// Listen and serve routes for router
	http.ListenAndServe(":8080", router)
}
