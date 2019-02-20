package routes

import (
	"net/http"

	"github.com/NullJupiter/GoTodoApp/src/front_end/templating"
)

// LoginGetHandler function is used to handle GET requests on /login
func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	templating.GetTemplates().ExecuteTemplate(w, "login", nil)
}

// LoginPostHandler function is used to handle POST requests on /login
func LoginPostHandler(w http.ResponseWriter, r *http.Request) {

}
