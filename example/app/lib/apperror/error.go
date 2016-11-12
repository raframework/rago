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
	typ  int
	code int
}

func New(typ int, code int) error {
	return &AppError{
		typ:  typ,
		code: code,
	}
}

func (ae *AppError) Error() string {
	return fmt.Sprintf("apperror type: %d, code: %d", ae.typ, ae.code)
}

func (ae *AppError) Typ() int {
	return ae.typ
}

func (ae *AppError) Code() int {
	return ae.code
}

func PanicWith(typ int, code int) {
	panic(New(typ, code))
}
