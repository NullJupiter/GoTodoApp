package queries

import (
	"fmt"
	"net/http"

	"github.com/NullJupiter/GoTodoApp/src/data_api/db"
	"github.com/NullJupiter/GoTodoApp/src/data_api/db/queries/helper"
)

// CreateUserEntry function is used to create a new user entry in the database.
// It also creates a new user specific table for the users todos.
func CreateUserEntry(username string, passwordHash string) error {
	// Check if user already exists
	uidCheck := helper.GetUIDForUname(username)
	if uidCheck != 0 {
		return fmt.Errorf("user already exists")
	}

	fmt.Println(1)

	// Create user entry in users table
	_, err := db.GetDB().Exec("INSERT INTO users (username, passwordHash) VALUES ($1, $2);", username, passwordHash)
	if err != nil {
		return err
	}
	fmt.Println(2)

	// Create user specific table for user
	uid := helper.GetUIDForUname(username)
	if uid == 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Println(3)

	_, err = db.GetDB().Exec(fmt.Sprintf("CREATE TABLE user_%v (id SERIAL PRIMARY KEY, todo varchar(255) NOT NULL, done integer NOT NULL);", uid))
	if err != nil {
		return err
	}
	fmt.Println(4)

	return nil
}
