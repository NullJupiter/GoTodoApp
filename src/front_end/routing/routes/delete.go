package routes

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/NullJupiter/GoTodoApp/src/front_end/cookiesessions"

	"github.com/gorilla/mux"
)

// DeleteGetHandler function is used to handle GET requests on /delete/{id}
func DeleteGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get Value from URL
	todoID := mux.Vars(r)["id"]

	// Get username from session cookie
	session, err := cookiesessions.GetCookieSessionStore().Get(r, "login-session")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// Extract username
	username := session.Values["username"].(string)
	// Save session
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Delete todo with API call
	resp, err := http.PostForm("http://localhost:8081/todos/delete", url.Values{"username": {username}, "todoID": {todoID}})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	respByteBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	respBody := string(respByteBody)
	if respBody == "Could not delete todo!" {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Redirect to /home
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
