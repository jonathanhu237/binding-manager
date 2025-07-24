package unierror

import "net/http"

var (
	ErrInternalServerError = &UnifiedError{
		Code:    http.StatusInternalServerError,
		Message: "The server encountered a problem and could not process your request.",
		Details: nil,
	}
)
