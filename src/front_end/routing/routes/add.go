package routes

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/NullJupiter/GoTodoApp/src/front_end/cookiesessions"
)

// AddPostHandler function is used to handle POST requests on /add
func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	// Get form values
	todo := r.FormValue("add-text")

	// Check if the form value is empty
	if todo == "" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

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

	// Add todo with API call on /todos/create
	resp, err := http.PostForm("http://localhost:8081/todos/create", url.Values{"username": {username}, "todo": {todo}, "done": {"0"}})
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
	if respBody != "Todo successfully inserted!" {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Redirect back to /home
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
