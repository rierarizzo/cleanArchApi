package error

import "errors"

var (
	ErrNotFound = errors.New("element not found")
	ErrUnknown  = errors.New("unknown error")
)
