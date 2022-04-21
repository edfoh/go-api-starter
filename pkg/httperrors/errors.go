package httperrors

import (
	"fmt"
)

type ClientErrorCode int

const (
	BadRequest ClientErrorCode = 400
)

func (code ClientErrorCode) New(msg string) error {
	return newClientError(int(code), msg)
}

func (code ClientErrorCode) Newf(msg string, args ...interface{}) error {
	return newClientError(int(code), fmt.Sprintf(msg, args...))
}

type clientError struct {
	code    int
	message string
}

func newClientError(code int, message string) *clientError {
	return &clientError{code, message}
}

func (e *clientError) Error() string {
	return fmt.Sprintf("statusCode: %d, message: %s", e.code, e.message)
}

func (e *clientError) ErrorDetails() (int, string) {
	return e.code, e.message
}
