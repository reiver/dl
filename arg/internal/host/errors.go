package arghost

import (
	"errors"
)

var (
	errNilDestination = errors.New("arghost: nil destination")
	errNotParsedYet   = errors.New("arghost: not parsed yet")
)
