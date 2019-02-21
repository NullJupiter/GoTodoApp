package queries

import (
	"github.com/NullJupiter/GoTodoApp/src/data_api/db"
	"github.com/NullJupiter/GoTodoApp/src/data_api/db/queries/helper"
	"golang.org/x/crypto/bcrypt"
)

// ValidateUser function is used to check if a user exists and returns true when it does (false when not).
func ValidateUser(username, password string) (bool, error) {
	uid := helper.GetUIDForUname(username)
	if uid == 0 {
		return false, nil
	}

	// Get passwordHash
	var passwordHash string
	err := db.GetDB().QueryRow("SELECT passwordHash FROM users WHERE id=$1", uid).Scan(&passwordHash)
	if err != nil {
		return false, err
	}

	// Compare password and passwordHash
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
