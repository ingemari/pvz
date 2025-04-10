package middleware

import (
	"os"
	"time"

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
