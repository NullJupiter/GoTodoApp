package routes

import (
	"net/http"

	"github.com/NullJupiter/GoTodoApp/src/front_end/templating"
)

// RegisterGetHandler function is used to handle GET requests on /register
func RegisterGetHandler(w http.ResponseWriter, r *http.Request) {
	templating.GetTemplates().ExecuteTemplate(w, "register", nil)
}

// RegisterPostHandler function is used to handle POST requests on /register
func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {

}
