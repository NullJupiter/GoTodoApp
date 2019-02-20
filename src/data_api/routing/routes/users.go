package routes

import (
	"net/http"

	"github.com/NullJupiter/GoTodoApp/src/data_api/db/queries"
)

// OneUserGetHandler function is used to handle GET requests on /user/{id}
func OneUserGetHandler(w http.ResponseWriter, r *http.Request) {

}

// CreateUserHandler function is used to handle POST requests in /user/create
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get username and passwordHash from POST data
	username := r.FormValue("username")
	passwordHash := r.FormValue("passwordHash")

	err := queries.CreateUserEntry(username, passwordHash)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
