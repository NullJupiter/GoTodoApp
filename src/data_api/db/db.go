package db

import (
	"database/sql"

	// Import postgres driver
	_ "github.com/lib/pq"
)

var err error
var db *sql.DB

// InitDB function is used to connect to a postgres database
func InitDB(dbURI string) error {
	db, err = sql.Open("postgres", dbURI)
	if err != nil {
		return err
	}

	return nil
}

// GetDB function is used to get the initialized database object from other packages
func GetDB() *sql.DB {
	return db
}
