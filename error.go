package mugard

import "errors"

var (
	ErrNilValue      = errors.New("value is nil")
	ErrMultipleWrite = errors.New("another process already get the write process")
)
