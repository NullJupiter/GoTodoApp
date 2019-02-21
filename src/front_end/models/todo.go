package models

// Todo type is used to convert json data from the API to a golang struct.
type Todo struct {
	ID   int    `json:"id"`
	Todo string `json:"todo"`
	Done int    `json:"done"`
}
