package model

import "errors"

// ErrUnknownType occurs if requested secret type is unknown.
var ErrUnknownType = errors.New("unknown type")
