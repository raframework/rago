package raerror

import (
	"fmt"
)

// ------------------------------------------------------------------

type BadBodyError struct {
	s string
}

func (b *BadBodyError) Error() string {
	return b.s
}

func NewBadBodyError(s string) *BadBodyError {
	return &BadBodyError{s}
}

// ------------------------------------------------------------------

type MethodNotAllowedError struct {
	s string
}

func (m *MethodNotAllowedError) Error() string {
	return m.s
}

func NewMethodNotAllowedError(s string) *MethodNotAllowedError {
	return &MethodNotAllowedError{s}
}

func NewMethodNotAllowedErrorf(format string, v ...interface{}) *MethodNotAllowedError {
	return &MethodNotAllowedError{fmt.Sprintf(format, v...)}
}

// ------------------------------------------------------------------

type NotFoundError struct {
	s string
}

func (n *NotFoundError) Error() string {
	return n.s
}

func NewNotFoundError(s string) *NotFoundError {
	return &NotFoundError{s}
}

func NewNotFoundErrorf(format string, v ...interface{}) *NotFoundError {
	return &NotFoundError{fmt.Sprintf(format, v...)}
}

// ------------------------------------------------------------------

type UnsupportedMediaTypeError struct {
	s string
}

func (u *UnsupportedMediaTypeError) Error() string {
	return u.s
}

func NewUnsupportedMediaTypeError(s string) *UnsupportedMediaTypeError {
	return &UnsupportedMediaTypeError{s}
}

// ------------------------------------------------------------------

type RuntimeError struct {
	s string
}

func (r *RuntimeError) Error() string {
	return r.s
}

func NewRuntimeError(s string) *RuntimeError {
	return &RuntimeError{s}
}

func NewRuntimeErrorf(format string, v ...interface{}) *RuntimeError {
	return &RuntimeError{fmt.Sprintf(format, v...)}
}
