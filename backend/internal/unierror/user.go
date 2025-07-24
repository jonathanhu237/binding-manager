package unierror

var (
	ErrUsernameAlreadyExists = &UnifiedError{
		Code:    100001,
		Message: "The username already exists.",
		Details: nil,
	}
	ErrEmailAlreadyExists = &UnifiedError{
		Code:    100002,
		Message: "The email already exists.",
		Details: nil,
	}
)
