package routes

import (
	"net/http"

	"github.com/NullJupiter/GoTodoApp/src/templating"
)

// MainGetHandler function is used to handle GET requests on /home
func MainGetHandler(w http.ResponseWriter, r *http.Request) {
	templating.GetTemplates().ExecuteTemplate(w, "todos", nil)
}
