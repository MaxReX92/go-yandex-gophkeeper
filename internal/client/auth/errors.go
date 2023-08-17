package auth

import "errors"

var (
	// ErrUnauthorized occurs if user auth credentials is missed.
	ErrUnauthorized = errors.New("unauthorized")
	// ErrAlreadyAuthorized occurs if no more authorization required.
	ErrAlreadyAuthorized = errors.New("already authorized")
)
