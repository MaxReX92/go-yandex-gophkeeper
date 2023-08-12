package crypto

import "errors"

var (
	ErrInvalidKey    = errors.New("invalid key")
	ErrTooShortBlock = errors.New("too short block")
)
