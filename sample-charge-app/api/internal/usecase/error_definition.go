package usecase

import "errors"

var (
	ErrInternal          = errors.New("internal server error")
	ErrNotFound          = errors.New("not found")
	ErrInvalidRequest    = errors.New("invalid request")
	ErrInvalidArgument   = errors.New("invalid argument")
	ErrInvalidState      = errors.New("invalid state")
	ErrInvalidTransition = errors.New("invalid transition")
	ErrInvalidOperation  = errors.New("invalid operation")
)
