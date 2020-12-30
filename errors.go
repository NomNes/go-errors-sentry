package errors

import (
	"errors"
	"runtime"
)

type Error struct {
	err error
	*stack
	extra Extra
	msg   string
}

type Extra = map[string]interface{}

func (e *Error) Error() string {
	msg := ""
	if e.err != nil {
		msg += e.err.Error()
	}
	if e.msg != "" {
		msg += " " + e.msg
	}
	return msg
}

func (e *Error) Extra() Extra {
	return e.extra
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

func Wrap(err error) *Error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*Error); ok {
		return e
	}
	return &Error{err, callers(), nil, ""}
}

func New(message string) error {
	return Wrap(errors.New(message))
}

func WrapExtra(err error, extra Extra) *Error {
	e := Wrap(err)
	if e != nil {
		e.extra = extra
	}
	return e
}

func NewExtra(message string, extra Extra) error {
	return WrapExtra(errors.New(message), extra)
}

func WrapMessage(err error, message string) *Error {
	e := Wrap(err)
	if e != nil {
		e.msg = message
	}
	return e
}
