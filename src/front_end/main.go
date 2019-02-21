package main

import (
	"log"
	"net/http"

	"github.com/NullJupiter/GoTodoApp/src/front_end/cookiesessions"
	"github.com/NullJupiter/GoTodoApp/src/front_end/templating"

	"github.com/NullJupiter/GoTodoApp/src/front_end/routing"
)

func main() {
	// initialize cookie session store
	err := cookiesessions.InitCookieSessionStore()
	if err != nil {
		log.Fatalf("Could not initialize cookie session store. Error: %v", err.Error())
	}

	// initialize template parsing
	err = templating.InitTemplating("views/templates/*.html", "views/codesnippets/*.html")
	if err != nil {
		log.Fatalf("Could not parse templates. Error: %v", err.Error())
	}

	// initialize routing
	router := routing.InitRouting()

	// serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// listen and serve on routes
	log.Fatal(http.ListenAndServe(":8080", router))
}
