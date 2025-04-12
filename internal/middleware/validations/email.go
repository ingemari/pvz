package validations

import (
	"errors"
	"regexp"
)

var ErrInvalidEmail = errors.New("invalid email")

func IsValidEmail(email string) bool {
	// Простая проверка через regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func IsValidRole(role string) bool {
	if role == "employee" || role == "moderator" {
		return true
	} else {
		return false
	}
}
