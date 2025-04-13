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

func IsValidPassword(p string) bool {
	if len(p) < 8 {
		return false
	}
	hasLetter := regexp.MustCompile(`[A-Za-z]`).MatchString(p)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(p)

	return hasLetter && hasNumber
}

func IsValidCity(city string) bool {
	if city == "Москва" || city == "Санкт-Петербург" || city == "Казань" {
		return true
	} else {
		return false
	}
}

func IsValidStatus(s string) bool {
	if s == "in_progress" || s == "close" {
		return true
	} else {
		return false
	}
}
