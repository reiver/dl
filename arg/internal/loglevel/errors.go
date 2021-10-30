package argloglevel

import (
	"errors"
)

var (
	errNilDestination = errors.New("argloglevel: nil destination")
	errNotParsedYet   = errors.New("argloglevel: not parsed yet")
)
