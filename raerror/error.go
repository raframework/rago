package raerror

import (
	"fmt"
)

type RaErrorTyp int

const (
	// System errors
	TypRuntime         RaErrorTyp = 1
	TypLogic                      = 2
	TypInvalidArgument            = 3

	// HTTP errors
	TypBadRequest           = 400
	TypNotFound             = 404
	TypMethodNotAllowed     = 405
	TypUnsupportedMediaType = 415
)

var errorTypStringMap = map[RaErrorTyp]string{
	TypRuntime:              "TypRuntime",
	TypLogic:                "TypLogic",
	TypInvalidArgument:      "TypInvalidArgument",
	TypBadRequest:           "TypBadRequest",
	TypNotFound:             "TypNotFound",
	TypMethodNotAllowed:     "TypMethodNotAllowed",
	TypUnsupportedMediaType: "TypUnsupportedMediaType",
}

// String returns a multi-character representation of the RaErrorTyp.
func (t RaErrorTyp) String() string {
	str, ok := errorTypStringMap[t]
	if !ok {
		panic(fmt.Sprintf("raerror: unknown RaErrorTyp %d", t))
	}

	return str
}

type RaError struct {
	typ     RaErrorTyp
	code    int
	message string
}

func New(typ RaErrorTyp, code int, message string) error {
	return &RaError{
		typ:     typ,
		code:    code,
		message: message,
	}
}

func (re *RaError) Error() string {
	return fmt.Sprintf("raerror: type(%s) code(%d) message(%s)", re.typ, re.code, re.message)
}

func (re *RaError) Typ() RaErrorTyp {
	return re.typ
}

func (re *RaError) Code() int {
	return re.code
}

func (re *RaError) Message() string {
	return re.message
}

func PanicWith(typ RaErrorTyp, code int, message string) {
	panic(New(typ, code, message))
}
