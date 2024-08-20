package utils

import (
	"strings"
	"unicode"
)

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsLength(s string, min int, max int) bool {
	temp := strings.TrimSpace(s)
	l := len(temp)
	return l >= min && l <= max
}

func Contains(validValues []string, val string) bool {
	for _, el := range validValues {
		if el == val {
			return true
		}
	}

	return false
}

func ValidPassword(s string) bool {
	temp := strings.TrimSpace(s)
	if len(temp) < 8 || len(temp) > 16 {
		return false
	}

	hasUpper, hasLower, hasDigit, hasSpecialChar := false, false, false, false
	allowedSpecialChar := "@!$&"

	for _, c := range temp {
		if unicode.IsUpper(c) {
			hasUpper = true
		} else if unicode.IsLower(c) {
			hasLower = true
		} else if unicode.IsDigit(c) {
			hasDigit = true
		} else if strings.ContainsRune(allowedSpecialChar, c) {
			hasSpecialChar = true
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecialChar
}
