package cli

import "errors"

var (
	ErrFileNotFound        = errors.New("file not found")
	ErrInvalidArguments    = errors.New("invalid arguments")
	ErrInvalidSecretType   = errors.New("invalid secret type")
	ErrRequiredArgNotFound = errors.New("required arg not found")
	ErrSecretNotFound      = errors.New("secret not found")
)
