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
	router.HandleFunc("/todos", routes.TodosPostHandler).Methods(http.MethodPost)
	router.HandleFunc("/todos/create", routes.CreateTodoPostHandler).Methods(http.MethodPost)
	router.HandleFunc("/todos/done", routes.DonePostHandler).Methods(http.MethodPost)
	router.HandleFunc("/todos/delete", routes.DeletePostHandler).Methods(http.MethodPost)
	router.HandleFunc("/user/validate", routes.ValidateUserGetHandler).Methods(http.MethodPost)
	router.HandleFunc("/user/create", routes.CreateUserHandler).Methods(http.MethodPost)

	return router
}
