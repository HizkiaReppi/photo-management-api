package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

const minCost = bcrypt.MinCost

// HashPassword generates a bcrypt hash for the provided password.
func HashPassword(password string) (string, error) {
	hashResult, err := bcrypt.GenerateFromPassword([]byte(password), minCost)
	if err != nil {
		return "", err
	}

	return string(hashResult), nil
}

// ComparePassword compares the input password with the stored hashed password.
func ComparePassword(inputPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
