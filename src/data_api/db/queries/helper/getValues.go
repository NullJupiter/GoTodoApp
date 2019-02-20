package helper

import (
	"database/sql"
	"fmt"

	"github.com/NullJupiter/GoTodoApp/src/data_api/db"
)

// GetUIDForUname function is used to query the id for a specific username
func GetUIDForUname(username string) interface{} {
	var uid interface{}
	err := db.DB.QueryRow("SELECT id FROM users WHERE username=$1;", username).Scan(&uid)
	fmt.Println(uid)
	if err == sql.ErrNoRows {
		return 0
	}
	if err != nil {
		return 0
	}

	return uid
}

// GetUIDForUnameAndPass function is used to query the id for a specific username password combination
func GetUIDForUnameAndPass(username, passwordHash string) (int, error) {
	row := db.DB.QueryRow("SELECT id FROM users WHERE username=$1 AND passwordHash=$2", username, passwordHash)
	var uid int
	err := row.Scan(&uid)
	if err != nil {
		return 0, err
	}

	return uid, nil
}
