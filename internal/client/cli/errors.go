package cli

import "errors"

var (
	ErrInvalidSecretType   = errors.New("invalid secret type")
	ErrRequiredArgNotFound = errors.New("required arg not found")
	ErrSecretNotFound      = errors.New("secret not found")
)
