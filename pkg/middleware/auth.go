package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store claims in context
		c.Set("claims", claims)
		c.Set("userID", claims.ID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("tokenType", claims.Type)

		c.Next()
	}
}

// AdminMiddleware checks if user has admin role
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userClaims := claims.(*Claims)
		if userClaims.Type != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// UserMiddleware checks if user has user role
func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userClaims := claims.(*Claims)
		if userClaims.Type != "user" {
			c.JSON(http.StatusForbidden, gin.H{"error": "User access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}
