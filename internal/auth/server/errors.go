package server

import "errors"

var (
	// ErrInvalidRequest occurs if invalid auth request was received.
	ErrInvalidRequest = errors.New("invalid request")

	// ErrInvalidCredentials occurs if invalid credentials was received in auth request.
	ErrInvalidCredentials = errors.New("invalid credentials")

	// ErrLoginNotFound occurs if provided credentials login was not found.
	ErrLoginNotFound = errors.New("login not found")
)
