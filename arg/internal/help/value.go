package arghelp

import (
	"flag"
)

var (
	value bool
)

func init() {
	flag.BoolVar(&value, "help", false, "output this help message")
}


func Receive(dst *bool) error {
	if !flag.Parsed() {
		return errNotParsedYet
	}

	if nil == dst {
		return errNilDestination
	}

	*dst = value
	return nil
}
