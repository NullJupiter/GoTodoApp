package routes

import (
	"fmt"
	"net/http"
)

// IndexGetHandler function is used to handle GET requests on /
func IndexGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Go to /todos, /user/validate or /user/create")
}
