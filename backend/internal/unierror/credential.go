package unierror

var (
	ErrInvalidCredentials = &UnifiedError{
		Code:    101001,
		Message: "Invalid credentials provided.",
		Details: nil,
	}
)
