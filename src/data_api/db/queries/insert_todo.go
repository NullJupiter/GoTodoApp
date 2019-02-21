package queries

import (
	"fmt"

	"github.com/NullJupiter/GoTodoApp/src/data_api/db"
	"github.com/NullJupiter/GoTodoApp/src/data_api/db/queries/helper"
)

// InsertTodoForUname function is used to insert a todo for a specific user
func InsertTodoForUname(username, todo string, done int) error {
	// Get uid for that user
	uid := helper.GetUIDForUname(username)

	// Insert todo into user table for specific user
	_, err := db.GetDB().Exec(fmt.Sprintf("INSERT INTO user_%v (todo, done) VALUES ($1, $2);", uid), todo, done)
	if err != nil {
		return err
	}

	return nil
}
