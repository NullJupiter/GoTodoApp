package models

// HomeData type is used to pass data to the home template (aka.: todo.html)
type HomeData struct {
	Username string
	Todos    []Todo
}
