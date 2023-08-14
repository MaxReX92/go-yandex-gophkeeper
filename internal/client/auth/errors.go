package auth

import "errors"

var (
	ErrUnauthorized      = errors.New("unauthorized")
	ErrAlreadyAuthorized = errors.New("already authorized")
)
