package routing

import (
	"net/http"

	"github.com/NullJupiter/GoTodoApp/src/front_end/routing/routes"
	"github.com/gorilla/mux"
)

// InitRouting function is used to initialize the routes and to return a http.Handler
func InitRouting() *mux.Router {
	router := mux.NewRouter()
	// index routing
	router.HandleFunc("/", routes.IndexGetHandler).Methods(http.MethodGet)
	// login routing
	router.HandleFunc("/login", routes.LoginGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", routes.LoginPostHandler).Methods(http.MethodPost)
	// register routing
	router.HandleFunc("/register", routes.RegisterGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/register", routes.RegisterPostHandler).Methods(http.MethodPost)
	// main page routing
	router.HandleFunc("/home", routes.MainGetHandler).Methods(http.MethodGet)
	// add post routing
	router.HandleFunc("/add", routes.AddPostHandler).Methods(http.MethodPost)

	return router
}
