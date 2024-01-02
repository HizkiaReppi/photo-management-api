package database

import (
	"fmt"
	"log"

	"rest-api/helpers"
	"rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// InitDB initializes the database connection.
func InitDB() {
	loadEnvironment()
	setupDatabase()
}

// MigrateDB migrates the database schema.
func MigrateDB() {
	stage := helpers.GetAsString("STAGE", "development")

	if stage == "development" || stage == "production" {
		db.Debug().AutoMigrate(&models.User{}, &models.Photo{})
	}
}

// GetDB returns the database connection.
func GetDB() *gorm.DB {
	return db
}

func loadEnvironment() {
	var path string
	stage := helpers.GetAsString("STAGE", "development")

	if stage == "testing" {
		path = ".env.testing"
	}
	if stage != "testing" {
		path = ".env"
	}

	helpers.LoadEnv(path)
}

func setupDatabase() {
	dbURI := getDatabaseURI()

	var err error
	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getDatabaseURI() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		helpers.GetAsString("DB_USER", "postgres"),
		helpers.GetAsString("DB_PASSWORD", "mysecretpassword"),
		helpers.GetAsString("DB_HOST", "localhost"),
		helpers.GetAsInt("DB_PORT", 5432),
		helpers.GetAsString("DB_NAME", "postgres"),
	)
}
