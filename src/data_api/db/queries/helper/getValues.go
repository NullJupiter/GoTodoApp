package helper

import (
	"github.com/NullJupiter/GoTodoApp/src/data_api/db"
)

var err error

// GetUIDForUname function is used to query the id for a specific username
func GetUIDForUname(username string) (int, error) {
	row := db.GetDB().QueryRow("SELECT id FROM users WHERE username=$1;", username)
	var uid int
	err = row.Scan(&uid)
	if err != nil {
		return 0, err
	}

	return uid, nil
}

// GetUIDForUnameAndPass function is used to query the id for a specific username password combination
func GetUIDForUnameAndPass(username, passwordHash string) (int, error) {
	row := db.GetDB().QueryRow("SELECT id FROM users WHERE username=$1 AND passwordHash=$2", username, passwordHash)
	var uid int
	err = row.Scan(&uid)
	if err != nil {
		return 0, err
	}

	return uid, nil
}
