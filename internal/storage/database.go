package storage

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", filepath.Join(os.Getenv("DATABASE_DIR")+"/"+os.Getenv("DB_NAME")))

	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	createTables()
}

func createTables() {
	dir, err := os.Open(os.Getenv("SCHEMA_DIR"))
	if err != nil {
		fmt.Println("Error opening directory: ", err)
	}
	defer dir.Close()

	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Println("Error reading directory: ", err)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".sql" {
			continue
		}

		filePath := filepath.Join(os.Getenv("SCHEMA_DIR"), file.Name())
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", filePath, err)
			continue
		}

		if !Up(string(data)) {
			fmt.Println("Failed to execute SQL from file:", filePath)
		}
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
