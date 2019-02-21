package routes

import (
	"net/http"

	"github.com/NullJupiter/GoTodoApp/src/data_api/db/queries"
)

// ValidateUserGetHandler function is used to handle POST requests on /user/validate
func ValidateUserGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get username and passwordHash from POST data
	username := r.FormValue("username")
	password := r.FormValue("password")

	isLegit, err := queries.ValidateUser(username, password)
	if !isLegit {
		w.Write([]byte("This user account doesn't exist!"))
		return
	}
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("This user account is a valid user account!"))
}

// CreateUserHandler function is used to handle POST requests in /user/create
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get username and passwordHash from POST data
	username := r.FormValue("username")
	passwordHash := r.FormValue("passwordHash")

	err := queries.CreateUserEntry(username, passwordHash)
	if err != nil {
		w.Write([]byte("Could not create user account because this username already exists."))
		return
	}

	w.Write([]byte("User has successfully been added."))
}
