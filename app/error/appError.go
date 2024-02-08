package error

import "errors"

var (
	ErrNotFound   = errors.New("element not found")
	ErrUnknown    = errors.New("unknown error")
	ErrConversion = errors.New("conversion error")
	ErrRepository = errors.New("repository error")
	ErrTimeout    = errors.New("timeout error")
)
