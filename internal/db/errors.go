package db

import "errors"

// ErrInvalidDBValue occurs if invalid value was received from db.
var ErrInvalidDBValue = errors.New("invalid db value")
