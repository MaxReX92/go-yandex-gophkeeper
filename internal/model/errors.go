package model

import "errors"

var (
	ErrInvalidType = errors.New("invalid type")
	ErrUnknownType = errors.New("unknown type")
)
