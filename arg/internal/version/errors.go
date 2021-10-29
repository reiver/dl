package argversion

import (
	"errors"
)

var (
	errNilDestination = errors.New("argversion: nil destination")
	errNotParsedYet   = errors.New("argversion: not parsed yet")
)
