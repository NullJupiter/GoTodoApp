package queries

import (
	"fmt"

	"github.com/NullJupiter/GoTodoApp/src/data_api/db"
	"github.com/NullJupiter/GoTodoApp/src/data_api/db/models"
	"github.com/NullJupiter/GoTodoApp/src/data_api/db/queries/helper"
)

// GetTodosForUname function is used to get all todos for a specific username.
func GetTodosForUname(username string) ([]models.Todo, error) {
	// Get uid for user
	uid := helper.GetUIDForUname(username)
	// Query todo rows
	rows, err := db.GetDB().Query(fmt.Sprintf("SELECT * FROM user_%v;", uid))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Fetch todos
	todos := make([]models.Todo, 0)
	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.ID, &todo.Todo, &todo.Done)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return todos, nil
}
