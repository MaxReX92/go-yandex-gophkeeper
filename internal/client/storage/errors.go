package storage

import "errors"

var (
	ErrSecretAlreadyExist = errors.New("secret already exist")
	ErrSecretNotFound     = errors.New("secret not found")
)
