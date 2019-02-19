package main

import (
	"log"
	"net/http"

	"github.com/NullJupiter/GoTodoApp/src/routing"
)

func main() {
	// initialize routing
	router := routing.InitRouting()
	// listen and serve on routes
	log.Fatal(http.ListenAndServe(":8080", router))
}
