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
	ErrUsernameNotExists = &UnifiedError{
		Code:    100003,
		Message: "The username does not exist.",
		Details: nil,
	}
)
