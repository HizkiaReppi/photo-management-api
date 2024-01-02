package main

import (
	"fmt"
	"log"
	"net/http"

	"rest-api/databases"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	// Initialize and migrate the database.
	initializeDatabase()

	http.HandleFunc("/", helloWorldHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initializeDatabase() {
	// Initialize the database.
	database.InitDB()

	// Migrate the database.
	database.MigrateDB()
}
