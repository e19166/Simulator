package db

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func InitDB() error {
	connStr := "user=postgres password=1234 dbname=simulator sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	return nil
}