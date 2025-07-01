package storage

import (
	"database/sql"
	"os"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("sqlite3", os.Getenv("DB_NAME"))

	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
}

func Up(stmt string) bool {
	_, err := DB.Exec(stmt)

	if err != nil {
		panic("Failed to create websites table: ")
		return false
	}

	return true
}
