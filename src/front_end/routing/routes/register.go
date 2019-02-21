package routes

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/sessions"

	"github.com/NullJupiter/GoTodoApp/src/front_end/cookiesessions"
	"golang.org/x/crypto/bcrypt"

	"github.com/NullJupiter/GoTodoApp/src/front_end/templating"
)

// RegisterGetHandler function is used to handle GET requests on /register
func RegisterGetHandler(w http.ResponseWriter, r *http.Request) {
	templating.GetTemplates().ExecuteTemplate(w, "register", nil)
}

// RegisterPostHandler function is used to handle POST requests on /register
func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	// Get POST form values
	username := r.FormValue("username")
	password := r.FormValue("pwd")

	if username == "" || password == "" {
		templating.GetTemplates().ExecuteTemplate(w, "register", "You need to type in a password and a username!")
	}

	// Hash password.
	passwordByteHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	passwordHash := string(passwordByteHash)

	// Create new user entry.
	resp, err := http.PostForm("http://localhost:8081/user/create", url.Values{"username": {username}, "passwordHash": {passwordHash}})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	respByteBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	respBody := string(respByteBody)

	// Check for error codes from the api
	if respBody != "User has successfully been added." {
		if respBody == "Could not create user account because this username already exists." {
			templating.GetTemplates().ExecuteTemplate(w, "register", "This username has already been taken!")
			return
		}
		return
	}

	// Create new session because no session exists.
	session, err := cookiesessions.GetCookieSessionStore().Get(r, "login-session")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// Set options for session.
	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: time.Now().Add(time.Hour).Second(),
	}
	// Set username in cookie.
	session.Values["username"] = username
	// Save session to client.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Redirect to /home with new cookie
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
