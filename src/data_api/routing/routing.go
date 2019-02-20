package routing

import (
	"net/http"

	"github.com/NullJupiter/GoTodoApp/src/data_api/routing/routes"
	"github.com/gorilla/mux"
)

// InitRouting function is used to initialize routing
func InitRouting() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", routes.IndexGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/todos", routes.TodosGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", routes.OneUserGetHandler).Methods(http.MethodGet)

	return router
}
