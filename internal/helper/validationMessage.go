package helper

import (
	"github.com/go-playground/validator/v10"
)

func ValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "Field is required"
	case "email":
		return "Field should be email"
	case "min":
		return "Field should be more than " + err.Param() + " characters"
	default:
		return "Invalid field format"
	}
}
