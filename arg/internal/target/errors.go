package argtarget

import (
	"errors"
)

var (
	errNilDestination = errors.New("argtarget: nil destination")
	errNotParsedYet   = errors.New("argtarget: not parsed yet")
)
