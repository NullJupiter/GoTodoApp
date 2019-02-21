package routing

import (
	"net/http"

	"github.com/NullJupiter/GoTodoApp/src/data_api/routing/routes"
	"github.com/gorilla/mux"
)

// InitRouting function is used to initialize routing
func InitRouting() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", routes.IndexGetHandler).Methods(http.MethodGet)                      // finished
	router.HandleFunc("/todos", routes.TodosPostHandler).Methods(http.MethodPost)               // finished
	router.HandleFunc("/user/validate", routes.ValidateUserGetHandler).Methods(http.MethodPost) // finished
	router.HandleFunc("/user/create", routes.CreateUserHandler).Methods(http.MethodPost)        // finished

	return router
}
