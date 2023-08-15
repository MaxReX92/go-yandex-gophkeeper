package crypto

import "errors"

var (
	// ErrTooShortBlock occurs if encryptor block size is too short.
	ErrTooShortBlock = errors.New("too short block")
)
