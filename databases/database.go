package database

import (
	"fmt"
	"log"

	"rest-api/helpers/env"
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
	stage := env.GetAsString("STAGE", "development")

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
	stage := env.GetAsString("STAGE", "development")

	if stage == "testing" {
		path = ".env.testing"
	}
	if stage != "testing" {
		path = ".env"
	}

	env.LoadEnv(path)
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
		env.GetAsString("DB_USER", "postgres"),
		env.GetAsString("DB_PASSWORD", "mysecretpassword"),
		env.GetAsString("DB_HOST", "localhost"),
		env.GetAsInt("DB_PORT", 5432),
		env.GetAsString("DB_NAME", "postgres"),
	)
}
