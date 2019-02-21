package routes

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/sessions"

	"github.com/NullJupiter/GoTodoApp/src/front_end/cookiesessions"

	"github.com/NullJupiter/GoTodoApp/src/front_end/templating"
)

// LoginGetHandler function is used to handle GET requests on /login
func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	templating.GetTemplates().ExecuteTemplate(w, "login", nil)
}

// LoginPostHandler function is used to handle POST requests on /login
func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	// Get POST form values
	username := r.FormValue("username")
	password := r.FormValue("pwd")

	if username == "" || password == "" {
		templating.GetTemplates().ExecuteTemplate(w, "login", "You need to insert a username and password!")
		return
	}

	// Hash password
	passwordByteHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	passwordHash := string(passwordByteHash)

	// Validate user account
	resp, err := http.PostForm("http://localhost:8081/user/validate", url.Values{"username": {username}, "passwordHash": {passwordHash}})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	// Read from response body
	respByteBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	respBody := string(respByteBody)
	// Check error codes from api
	if respBody != "This user account is a valid user account!" {
		if respBody == "This user account doesn't exist!" {
			templating.GetTemplates().ExecuteTemplate(w, "login", "This account doesn't exist!")
			return
		}
		return
	}

	// Create a new session because there is none
	session, err := cookiesessions.GetCookieSessionStore().Get(r, "login-session")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// Set session options
	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: time.Now().Add(time.Hour).Second(),
	}
	// Set username for session
	session.Values["username"] = username
	// Save session
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// If the user is validated and the login-session cookie was saved redirect to /home
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
