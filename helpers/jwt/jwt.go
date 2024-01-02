package jwt

import (
	"errors"

	"rest-api/helpers/env"
	"github.com/dgrijalva/jwt-go"
)

// SecretKey merupakan kunci rahasia untuk penandatanganan token JWT.
var SecretKey = []byte(env.GetAsString("STAGE", "kuncirahasia"))

// GenerateToken digunakan untuk membuat token JWT.
func GenerateToken(userID int) (string, error) {
	// Set payload
	claims := jwt.MapClaims{"user_id": userID}

	// Set algorithm dan buat token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token menggunakan kunci rahasia
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken digunakan untuk memvalidasi dan mengembalikan token JWT yang sudah divalidasi.
func ValidateToken(encodedToken string) (*jwt.Token, error) {
	// Parse token dengan kunci rahasia
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		// Validasi metode penandatanganan
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		// Kembalikan kunci rahasia
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
