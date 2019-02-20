package queries

import (
	"fmt"

	"github.com/NullJupiter/GoTodoApp/src/data_api/db"
	"github.com/NullJupiter/GoTodoApp/src/data_api/db/queries/helper"
)

var err error

// CreateUserEntry function is used to create a new user entry in the database.
// It also creates a new user specific table for the users todos.
func CreateUserEntry(username string, passwordHash string) error {
	// Check if user already exists
	uidCheck, err := helper.GetUIDForUname(username)
	if uidCheck != 0 {
		return fmt.Errorf("user already exists")
	}

	// Create user entry in users table
	_, err = db.GetDB().Exec("INSERT INTO users (username, passwordHash) VALUES ($1, $2);", username, passwordHash)
	if err != nil {
		return err
	}

	// Create user specific table for user
	uid, err := helper.GetUIDForUname(username)
	if err != nil {
		return err
	}
	fmt.Println(1)

	_, err = db.GetDB().Exec(fmt.Sprintf("CREATE TABLE user_%v (id SERIAL PRIMARY KEY, todo varchar(255) NOT NULL, done integer NOT NULL);", uid))
	if err != nil {
		return err
	}
	fmt.Println(2)

	return nil
}
