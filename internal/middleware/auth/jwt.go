package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(role string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	claims := jwt.MapClaims{
		"role": role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func RequireRole(required string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsAny, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Access denied"})
			return
		}

		claims := claimsAny.(jwt.MapClaims)
		userRole, ok := claims["role"].(string)
		if !ok || userRole != required {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Access denied"})
			return
		}

		c.Next()
	}
}

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing Authorization header"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
