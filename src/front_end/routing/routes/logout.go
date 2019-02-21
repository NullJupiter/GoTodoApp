package routes

import (
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/NullJupiter/GoTodoApp/src/front_end/cookiesessions"
)

// LogoutGetHandler function is used to handle GET requests on /logout.
func LogoutGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get current cookie
	session, err := cookiesessions.GetCookieSessionStore().Get(r, "login-session")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// Let the cookie expire right now
	session.Options = &sessions.Options{
		MaxAge: -1,
	}
	// Save new immediatelly expiring cookie
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Redirect to /login
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
