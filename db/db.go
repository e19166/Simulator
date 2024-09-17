package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import PostgreSQL driver
)

// DB is a global variable that holds the database connection pool
var DB *sql.DB

// InitDB initializes the database connection
// It connects to the PostgreSQL database using the provided connection string
func InitDB() error {
	// Connection string to the PostgreSQL database
	connStr := "user=postgres password=1234 dbname=event_kind_db sslmode=disable"

	var err error
	// Open a database connection using the PostgreSQL driver
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		// Return an error if opening the database connection fails
		return fmt.Errorf("error opening database: %w", err)
	}

	// Ping the database to ensure the connection is established
	err = DB.Ping()
	if err != nil {
		// Return an error if connecting to the database fails
		return fmt.Errorf("error connecting to database: %w", err)
	}

	// Print a message indicating that the database connection was established successfully
	fmt.Println("Database connection established")
	return nil
}

// CloseDB closes the database connection
// It is important to close the connection to release resources
func CloseDB() {
	if DB != nil {
		// Attempt to close the database connection
		err := DB.Close()
		if err != nil {
			// Print an error message if closing the connection fails
			fmt.Printf("error closing database: %v\n", err)
		} else {
			// Print a message indicating that the database connection was closed successfully
			fmt.Println("Database connection closed")
		}
	}
}
