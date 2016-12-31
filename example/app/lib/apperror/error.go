package apperror

import (
	"fmt"
)

const (
	BadRequest          = 400
	Forbidden           = 403
	NotFound            = 404
	Unauthorized        = 401
	UnprocessableEntity = 422

	InternalServerError = 500
	BadGateway          = 502
)

type AppError struct {
	typ     int
	code    int
	message string
}

func New(typ int, code int, message string) error {
	return &AppError{
		typ:     typ,
		code:    code,
		message: message,
	}
}

func (ae *AppError) Error() string {
	return fmt.Sprintf("apperror: type(%d) code(%d) message(%s)", ae.typ, ae.code, ae.message)
}

func (ae *AppError) Typ() int {
	return ae.typ
}

func (ae *AppError) Code() int {
	return ae.code
}

func (ae *AppError) Message() string {
	return ae.message
}
