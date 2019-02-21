package routes

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/NullJupiter/GoTodoApp/src/front_end/cookiesessions"

	"github.com/gorilla/mux"
)

// DoneGetHandler function is used to handle GET requests on /done/{id}
func DoneGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get values from url
	todoID := mux.Vars(r)["id"]
	done := r.URL.Query()["finished"][0]

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

	// Change done value with API call
	resp, err := http.PostForm("http://localhost:8081/todos/done", url.Values{"username": {username}, "todoID": {todoID}, "done": {done}})
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
	if respBody == "Could not change done value of todo!" {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Redirect back to /home
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
