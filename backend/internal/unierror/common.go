package unierror

import (
	"errors"
	"net/http"

	"github.com/jonathanhu237/binding-manager/backend/internal/validator"
)

var (
	ErrInternalServerError = &UnifiedError{
		Code:    http.StatusInternalServerError,
		Message: "The server encountered a problem and could not process your request.",
		Details: nil,
	}
)

func ErrBadRequest(err error) *UnifiedError {
	unifiedErr := &UnifiedError{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
		Details: nil,
	}

	var validationErrs *validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		details := make(map[string]any)

		for key, msg := range *validationErrs {
			details[key] = msg
		}

		unifiedErr.Message = "Validation failed."
		unifiedErr.Details = &details
	}

	return unifiedErr
}
