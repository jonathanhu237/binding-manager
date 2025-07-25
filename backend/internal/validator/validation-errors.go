package validator

import "fmt"

type ValidationErrors map[string]string

func (e *ValidationErrors) Error() string {
	if len(*e) == 0 {
		return "No validation errors."
	}

	var errMsg string
	for key, message := range *e {
		if errMsg != "" {
			errMsg += "; "
		}
		errMsg += fmt.Sprintf("%s: %s", key, message)
	}

	return errMsg
}
