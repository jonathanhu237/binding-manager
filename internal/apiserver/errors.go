package apiserver

import "net/http"

type ServiceError struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Details *envelope `json:"details"`
}

var (
	ErrInternalServerError = ServiceError{
		Code:    http.StatusInternalServerError,
		Message: "The server encountered a problem and could not process your request.",
		Details: nil,
	}
)
