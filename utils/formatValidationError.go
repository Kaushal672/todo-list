package utils

import "github.com/go-playground/validator/v10"

func FormatValidationError(errs validator.ValidationErrors) map[string]string {
	errMessages := make(map[string]string)

	for _, e := range errs {
		switch e.Tag() {
		case "required":
			errMessages[e.Field()] = e.Field() + " is required"
		case "min":
			errMessages[e.Field()] = e.Field() + " must have atleast " + e.Param() + " characters"
		case "max":
			errMessages[e.Field()] = e.Field() + " must have atmost " + e.Param() + " characters"
		case "oneof":
			errMessages[e.Field()] = e.Field() + " should be one of the following: " + e.Param()
		case "password":
			errMessages[e.Field()] = "Password must include characters between 8 and 16, one uppercase, one lowercase, one digit and one special character"
		default:
			errMessages[e.Field()] = e.Field() + " is invalid"
		}
	}

	return errMessages
}
