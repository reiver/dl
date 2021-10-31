package optint

import (
	"errors"
)

var (
	errNilReceiver = errors.New("optint: nil receiver")
	errNilSource   = errors.New("optint: nil source")
	errNothing     = errors.New("optint: nothing")
)
