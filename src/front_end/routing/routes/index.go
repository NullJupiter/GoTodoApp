package routes

import "net/http"

// IndexGetHandler function is used to handle GET requests on /
func IndexGetHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
