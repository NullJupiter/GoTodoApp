package routes

import (
	"net/http"
)

// IndexGetHandler function is used to handle GET requests on /
func IndexGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Go to /todos/{id} or /users/{id}"))
}
