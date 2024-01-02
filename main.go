package main

import (
	"github.com/gin-gonic/gin"

	"rest-api/databases"
	"rest-api/routers"
)

func main() {
	// Initialize and migrate the database.
	initializeDatabase()

	// Initialize and run the router.
	initializeRouter().Run(":8080")
}

func initializeDatabase() {
	// Initialize the database.
	database.InitDB()

	// Migrate the database.
	database.MigrateDB()
}

func initializeRouter() *gin.Engine {
	// Initialize the router.
	return routers.RouteInit()
}
