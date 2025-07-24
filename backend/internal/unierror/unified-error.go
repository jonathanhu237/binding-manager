package unierror

type UnifiedError struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Details *map[string]any `json:"details"`
}

func (e *UnifiedError) Error() string {
	return e.Message
}
