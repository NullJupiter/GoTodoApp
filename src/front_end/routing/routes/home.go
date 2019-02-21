package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/sessions"

	"github.com/NullJupiter/GoTodoApp/src/front_end/models"

	"github.com/NullJupiter/GoTodoApp/src/front_end/cookiesessions"

	"github.com/NullJupiter/GoTodoApp/src/front_end/templating"
)

// MainGetHandler function is used to handle GET requests on /home
func MainGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get cookie from client.
	session, err := cookiesessions.GetCookieSessionStore().Get(r, "login-session")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Check if session already existed else redirect
	if session.IsNew {
		session.Options = &sessions.Options{
			Path:   "/",
			MaxAge: -1,
		}
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Get username from cookie store
	username := session.Values["username"].(string)
	// Save the sassion
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Fetch todos for that username from API
	resp, err := http.PostForm("http://localhost:8081/todos", url.Values{"username": {username}})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read from response
	respByteBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Convert response json to a slice of type struct Todo
	var todos []models.Todo
	err = json.Unmarshal(respByteBody, &todos)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Create HomeData struct
	var homeData models.HomeData
	homeData.Username = username
	homeData.Todos = todos

	templating.GetTemplates().ExecuteTemplate(w, "todos", homeData)
}
