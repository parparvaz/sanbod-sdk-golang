package sanbod

import (
	"errors"
	"fmt"
)

func (e APIError) Error() string {
	return fmt.Sprintf("<APIError> code=%d, msg=%s", e.ResultNumber, e.Message)
}

func IsAPIError(e error) bool {
	var aPIError *APIError
	ok := errors.As(e, &aPIError)
	return ok
}

type APIError struct {
	Err     bool `json:"error"`
	Message struct {
		Scope []string `json:"scope"`
	} `json:"message"`
	ResultNumber int64 `json:"result_number"`
}
