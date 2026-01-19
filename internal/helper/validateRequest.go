package helper

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateRequest(request any) map[string]string {
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		var validationError validator.ValidationErrors

		if errors.As(err, &validationError) {
			errors := make(map[string]string)

			for _, err := range validationError {
				field := strings.ToLower(err.Field())
				errors[field] = ValidationMessage(err)
			}

			return errors
		}
	}

	return nil
}
