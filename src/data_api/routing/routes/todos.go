package routes

import (
	"encoding/json"
	"net/http"

	"github.com/NullJupiter/GoTodoApp/src/data_api/db/queries"
)

// TodosPostHandler function is used to handle GET requests on /todos
func TodosPostHandler(w http.ResponseWriter, r *http.Request) {
	// Get username value from POST form
	username := r.FormValue("username")

	// Get the todos for the username
	todos, err := queries.GetTodosForUname(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Encode todos as json and send it as response
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
