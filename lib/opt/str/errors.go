package optstr

import (
	"errors"
)

var (
	errNilReceiver = errors.New("optstr: nil receiver")
	errNilSource   = errors.New("optstr: nil source")
	errNothing     = errors.New("optstr: nothing")
)
