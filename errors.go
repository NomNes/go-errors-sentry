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

func (e *Error) Err() error {
	return e.err
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

func Wrap(err error) error {
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

func WrapExtra(err error, extra Extra) error {
	e := Wrap(err)
	if e != nil {
		e.(*Error).extra = extra
	}
	return e
}

func NewExtra(message string, extra Extra) error {
	return WrapExtra(errors.New(message), extra)
}

func WrapMessage(err error, message string) error {
	e := Wrap(err)
	if e != nil {
		e.(*Error).msg = message
	}
	return e
}
