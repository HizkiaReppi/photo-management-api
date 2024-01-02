package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	jwtHelpers"rest-api/helpers/jwt"
	"rest-api/helpers/formatter"
	"rest-api/models"
	"gorm.io/gorm"
)

// AuthMiddleware berfungsi sebagai middleware otentikasi.
func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			handleUnauthorized(c)
			return
		}

		tokenString := getTokenFromHeader(authHeader)
		token, err := jwtHelpers.ValidateToken(tokenString)
		if err != nil {
			handleUnauthorized(c)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			handleUnauthorized(c)
			return
		}

		userID, err := getUserIDFromClaims(claims)
		if err != nil {
			handleUnauthorized(c)
			return
		}

		user, err := getUserByID(db, userID)
		if err != nil {
			handleUnauthorized(c)
			return
		}

		c.Set("currentUser", user)
	}
}

// handleUnauthorized menangani respons untuk permintaan yang tidak diotorisasi.
func handleUnauthorized(c *gin.Context) {
	response := formatter.ApiResponse(http.StatusUnauthorized, "error", nil, "Unauthorized")
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// getTokenFromHeader mendapatkan token dari header Authorization.
func getTokenFromHeader(authHeader string) string {
	dataToken := strings.Split(authHeader, " ")
	if len(dataToken) == 2 {
		return dataToken[1]
	}
	return ""
}

// getUserIDFromClaims mendapatkan user ID dari JWT claims.
func getUserIDFromClaims(claims jwt.MapClaims) (int, error) {
	userID := claims["user_id"].(float64)
	return int(userID), nil
}

// getUserByID mendapatkan user dari database berdasarkan ID.
func getUserByID(db *gorm.DB, userID int) (*models.User, error) {
	var user models.User
	err := db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
