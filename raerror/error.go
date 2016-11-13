package raerror

import (
	"fmt"
)

const (
	TypLogic           = 1
	TypInvalidArgument = 2
	TypRuntime         = 3

	TypNotFound             = 11
	TypMethodNotAllowed     = 12
	TypUnsupportedMediaType = 13
	TypBadBody              = 14
)

type RaError struct {
	typ     int
	code    int
	message string
}

func New(typ int, code int, message string) error {
	return &RaError{
		typ:     typ,
		code:    code,
		message: message,
	}
}

func (re *RaError) Error() string {
	return fmt.Sprintf("raerror type: %d, code: %d", re.typ, re.code)
}

func (re *RaError) Typ() int {
	return re.typ
}

func (re *RaError) Code() int {
	return re.code
}

func (re *RaError) Message() string {
	return re.message
}

func PanicWith(typ int, code int, message string) {
	panic(New(typ, code, message))
}
