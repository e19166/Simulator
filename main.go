package main

import (
	"log"

	"github.com/e19166/Simulator/db"
	"github.com/e19166/Simulator/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	err := db.InitDB()
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
	}
	defer db.CloseDB()

	// Create a new Gin router
	r := gin.Default()

	// Register the handler function for the '/projects' endpoint
	r.POST("/projects", handlers.CreateProject)

	// Start the HTTP server on port 8080
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}