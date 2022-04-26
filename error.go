package tiktok

import (
	"errors"
	"fmt"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

// ErrAppInfoEmpty means something is Missing.
var ErrAppInfoEmpty = Error("AppKey or AppSecret is empty")

type APIError struct {
	Code      int
	Message   string
	RequestID string
}

func (a *APIError) Error() string {
	return fmt.Sprintf("%d: %s", a.Code, a.Message)
}

func ErrCode(err error) int {
	var e *APIError
	if errors.As(err, &e) {
		return e.Code
	}
	return 0
}
