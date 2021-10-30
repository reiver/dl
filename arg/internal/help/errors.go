package arghelp

import (
	"errors"
)

var (
	errNilDestination = errors.New("arghelp: nil destination")
	errNotParsedYet   = errors.New("arghelp: not parsed yet")
)
