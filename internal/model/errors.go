package model

import "errors"

var (
	// ErrUnknownType occurs if requested secret type is unknown.
	ErrUnknownType = errors.New("unknown type")
)
