package queries

import (
	"fmt"

	"github.com/NullJupiter/GoTodoApp/src/data_api/db"
)

var err error

// CreateUserEntry function is used to create a new user entry in the database.
// It also creates a new user specific table for the users todos.
func CreateUserEntry(username string, passwordHash string) error {
	// Create user entry in users table
	_, err = db.GetDB().Exec("INSERT INTO users (username, passwordHash) VALUES ($1, $2);", username, passwordHash)
	if err != nil {
		return err
	}

	// Create user specific table for user
	uid, err := GetUIDForUname(username)
	if err != nil {
		return err
	}
	_, err = db.GetDB().Exec(fmt.Sprintf("CREATE TABLE user_%v (id SERIAL PRIMARY KEY, todo varchar(255) NOT NULL, done integer;", uid))
	if err != nil {
		return err
	}

	return nil
}
