package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key-change-this-in-production")

type Claims struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"` // "user", "admin", "super_admin"
	Type     string `json:"type"` // "user" or "admin"
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT token
func GenerateToken(id int, email string, username string, role string, tokenType string) (string, error) {
	claims := &Claims{
		ID:       id,
		Email:    email,
		Username: username,
		Role:     role,
		Type:     tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken verifies and parses the JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// SetJWTSecret allows changing the secret key
func SetJWTSecret(secret string) {
	jwtSecret = []byte(secret)
}
