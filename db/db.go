package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import PostgreSQL driver
)

// DB is a global variable that holds the database connection pool
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	connStr := "user=postgres password=1234 dbname=event_kind_db sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	fmt.Println("Database connection established")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			fmt.Printf("error closing database: %v\n", err)
		} else {
			fmt.Println("Database connection closed")
		}
	}
}