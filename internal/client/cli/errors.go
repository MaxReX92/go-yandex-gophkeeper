package cli

import "errors"

var (
	// ErrFileNotFound occurs if binary secret file was not found.
	ErrFileNotFound = errors.New("file not found")
	// ErrInvalidArguments occurs if passed arguments list is invalid.
	ErrInvalidArguments = errors.New("invalid arguments")
	// ErrInvalidSecretType occurs if secret type is invalid.
	ErrInvalidSecretType = errors.New("invalid secret type")
	// ErrRequiredArgNotFound occurs if required command arg was not found.
	ErrRequiredArgNotFound = errors.New("required arg not found")
	// ErrSecretNotFound occurs if requestd secret was not found.
	ErrSecretNotFound = errors.New("secret not found")
)
