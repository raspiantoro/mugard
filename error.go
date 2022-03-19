package mugard

import "errors"

var (
	ErrNilValue              = errors.New("value is nil")
	ErrMultipleWrite         = errors.New("another process already get the write process")
	ErrNoLockedResources     = errors.New("there is no locked resources")
	ErrNotHoldingWriteAccess = errors.New("not holding the write access")
)
