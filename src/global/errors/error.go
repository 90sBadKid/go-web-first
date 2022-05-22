package errors

import (
	"go-web/global/result"
)

type ServerError struct {
	code    int
	message string
}

func NewServerError(code int, message string) *ServerError {
	return &ServerError{
		code:    code,
		message: message,
	}
}

func (e *ServerError) Code() int {
	return e.code
}

func (e *ServerError) Error() string {
	return result.StatusText(result.StatusServerError)
}
