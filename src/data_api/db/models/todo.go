package models

// Todo type is used to pass on todos from the database.
type Todo struct {
	ID   int    `json:"id"`
	Todo string `json:"todo"`
	Done int    `json:"done"`
}
