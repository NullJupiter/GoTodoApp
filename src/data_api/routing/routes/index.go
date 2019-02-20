package routes

import (
	"fmt"
	"net/http"
)

// IndexGetRouter function is used to handle GET requests on /
func IndexGetRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Go to /todos/{id} or /users/{id}")
}
