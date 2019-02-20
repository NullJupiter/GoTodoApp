package routes

import (
	"net/http"
)

// IndexGetRouter function is used to handle GET requests on /
func IndexGetRouter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Go to /todos/{id} or /users/{id}"))
}
