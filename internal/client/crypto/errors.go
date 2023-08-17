package crypto

import "errors"

// ErrTooShortBlock occurs if encryptor block size is too short.
var ErrTooShortBlock = errors.New("too short block")
