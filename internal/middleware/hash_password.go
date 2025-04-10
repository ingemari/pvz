package middleware

import (
	"fmt"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Failed to hash pass")
		return nil, fmt.Errorf("не удалось хешировать пароль: %w", err)
	}

	return hashedPassword, nil
}
