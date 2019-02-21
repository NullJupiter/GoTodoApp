package queries

import (
	"fmt"

	"github.com/NullJupiter/GoTodoApp/src/data_api/db"
	"github.com/NullJupiter/GoTodoApp/src/data_api/db/queries/helper"
)

// ChangeDoneValue function is used to change the done value of a specific todo.
func ChangeDoneValue(username string, todoID, done int) error {
	uid := helper.GetUIDForUname(username)

	_, err := db.GetDB().Exec(fmt.Sprintf("UPDATE user_%v SET done=$1 WHERE id=$2;", uid), done, todoID)
	if err != nil {
		return err
	}

	return nil
}
