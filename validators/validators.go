package validators

import (
	"regexp"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	// log.Println("init function invoked")
	Validate = validator.New(validator.WithRequiredStructEnabled())
	Validate.RegisterValidation("password", ValidatePassword)
}

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 || len(password) > 16 {
		return false
	}

	hasUpper, hasLower, hasDigit := false, false, false
	hasSpecialChar := regexp.MustCompile(`[!@#~$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+`).
		MatchString(password)

	for _, c := range password {
		if unicode.IsUpper(c) {
			hasUpper = true
		} else if unicode.IsLower(c) {
			hasLower = true
		} else if unicode.IsDigit(c) {
			hasDigit = true
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecialChar
}
