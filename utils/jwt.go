package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// JWT secret key
var jwtSecret = []byte("your_secret_key") // Replace with a more secure key in production

// GenerateToken generates a new JWT token
func GenerateToken(email, role string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // Token expiration time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken validates the JWT token and returns claims
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Parse token dengan callback untuk memverifikasi metode signing
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan metode signing adalah HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	// Jika terjadi error atau token tidak valid, langsung return error
	if err != nil || token == nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	// Konversi claims ke jwt.MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unable to parse claims")
	}

	return claims, nil
}
