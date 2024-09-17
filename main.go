package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/e19166/Simulator/db"
	"github.com/e19166/Simulator/handlers"
)

func main() {
	// Initiate database connection
	err := db.InitDB()
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
	}
	defer db.DB.Close()

	// Create a new router
	r := mux.NewRouter()

	// Register the handler functions
	r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}