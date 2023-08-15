package storage

import "errors"

var (
	// ErrSecretAlreadyExist occurs if requested secret is already exists.
	ErrSecretAlreadyExist = errors.New("secret already exist")
	// ErrSecretNotFound occurs if requested secret was not found.
	ErrSecretNotFound = errors.New("secret not found")
)
