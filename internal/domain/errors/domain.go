package errors

import "errors"

var (
	ErrNotFound          = errors.New("not found")
	ErrBadRequest        = errors.New("bad request")
	ErrInternal          = errors.New("internal server error")
	ErrForbidden         = errors.New("forbidden")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrExpectationFailed = errors.New("expectation failed")
	ErrInvalidAppKey     = errors.New("invalid app key")
)
