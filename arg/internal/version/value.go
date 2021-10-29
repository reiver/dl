package argversion

import (
	"flag"
)

var (
	value bool
)

func init() {
	flag.BoolVar(&value, "version", false, "output the version information and exit")
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
