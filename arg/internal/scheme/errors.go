package argscheme

import (
	"errors"
)

var (
	errNilDestination = errors.New("argscheme: nil destination")
	errNotParsedYet   = errors.New("argscheme: not parsed yet")
)
