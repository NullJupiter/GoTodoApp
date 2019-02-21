package queries

import (
	"fmt"

	"github.com/NullJupiter/GoTodoApp/src/data_api/db"
	"github.com/NullJupiter/GoTodoApp/src/data_api/db/queries/helper"
)

// DeleteTodo function is used to delete a specific todo for a specific user.
func DeleteTodo(username string, todoID int) error {
	uid := helper.GetUIDForUname(username)

	_, err := db.GetDB().Exec(fmt.Sprintf("DELETE FROM user_%v WHERE id=$1", uid), todoID)
	if err != nil {
		return err
	}

	return nil
}
