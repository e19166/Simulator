package main

import (
	"log"
	"net/http"

	"github.com/e19166/Simulator/db"
	"github.com/e19166/Simulator/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	// This function sets up the connection to the PostgreSQL database
	err := db.InitDB()
	if err != nil {
		// Log and terminate if there's an error initializing the database
		log.Fatalf("error initializing database: %v", err)
	}
	// Ensure the database connection is closed when the application exits
	defer db.DB.Close()

	// Create a new router using Gorilla Mux
	r := mux.NewRouter()

	// Register the handler function for the '/projects' endpoint
	// This sets up the route to handle POST requests to create new projects
	r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")

	// Start the HTTP server on port 8080
	// This listens for incoming requests and directs them to the router
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
