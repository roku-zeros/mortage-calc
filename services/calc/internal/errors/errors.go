package pkgerrors

import "errors"

const (
	ErrMsgNoProgram          = "choose program"
	ErrMsgMoreThanOneProgram = "choose only 1 program"
	ErrMsgBadInitialPayment  = "the initial payment should be more"
	ErrMsgEmptyCache         = "empty cache"
)

var (
	ErrNoProgram          = errors.New(ErrMsgNoProgram)
	ErrMoreThanOneProgram = errors.New(ErrMsgMoreThanOneProgram)
	ErrBadInitialPayment  = errors.New(ErrMsgBadInitialPayment)
	ErrEmptyCache         = errors.New(ErrMsgEmptyCache)
)
