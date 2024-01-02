package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the specified file path.
func LoadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetAsString retrieves the value of the environment variable as a string.
// If the variable is not set, it returns the specified defaultValue.
func GetAsString(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetAsInt retrieves the value of the environment variable as an integer.
// If the variable is not set or conversion fails, it returns the specified defaultValue.
func GetAsInt(name string, defaultValue int) int {
	valueStr := GetAsString(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
