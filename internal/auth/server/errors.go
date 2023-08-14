package server

import "errors"

var (
	ErrInvalidRequest     = errors.New("invalid request")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrLoginNotFound      = errors.New("login not found")
)
