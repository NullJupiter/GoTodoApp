package helper

import (
	"crypto/rand"

	"github.com/NullJupiter/GoTodoApp/src/front_end/models"
)

// GenerateRandomByteSlice function is used to generate a truly random byte slice for signing the cookies.
func GenerateRandomByteSlice(size int) ([]byte, error) {
	byteSlice := make([]byte, size)
	_, err := rand.Read(byteSlice)
	if err != nil {
		return nil, err
	}

	return byteSlice, nil
}

// ReverseTodoSlice function is used to reverse a slice todos.
func ReverseTodoSlice(slice []models.Todo) []models.Todo {
	outputSlice := make([]models.Todo, len(slice))
	for i := 0; i < len(slice); i++ {
		outputSlice[i] = slice[len(slice)-1-i]
	}

	return outputSlice
}
