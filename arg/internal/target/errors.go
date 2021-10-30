package argtarget

import (
	"errors"
)

var (
	errNilDestination = errors.New("argtarget: nil destination")
	errNotFound       = errors.New("argtarget: not found")
	errNotParsedYet   = errors.New("argtarget: not parsed yet")
)
