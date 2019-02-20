package queries

import "github.com/NullJupiter/GoTodoApp/src/data_api/db/queries/helper"

// ValidateUser function is used to check if a user exists and returns true when it does (false when not).
func ValidateUser(username, passwordHash string) (bool, error) {
	uid, err := helper.GetUIDForUnameAndPass(username, passwordHash)
	if err != nil {
		return false, err
	}
	if uid == 0 {
		return false, nil
	}

	return true, nil
}
