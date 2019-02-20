package helper

import (
	"context"
	"database/sql"

	"github.com/NullJupiter/GoTodoApp/src/data_api/db"
)

var ctx = context.Background()

// GetUIDForUname function is used to query the id for a specific username
func GetUIDForUname(username string) string {
	var uid string
	var uname string
	var pass string
	err := db.GetDB().QueryRow("SELECT * FROM users WHERE username='NullJupiter;").Scan(&uid, &uname, &pass)
	if err == sql.ErrNoRows {
		return "0"
	}
	if err != nil {
		return "0"
	}

	return uid
}

// GetUIDForUnameAndPass function is used to query the id for a specific username password combination
func GetUIDForUnameAndPass(username, passwordHash string) (int, error) {
	row := db.GetDB().QueryRow("SELECT id FROM users WHERE username=$1 AND passwordHash=$2", username, passwordHash)
	var uid int
	err := row.Scan(&uid)
	if err != nil {
		return 0, err
	}

	return uid, nil
}
