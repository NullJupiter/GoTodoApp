package routes

import (
	"fmt"
	"net/http"
)

func IndexGetRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Go to /todos/{id} or /users/{id}")
}
