package argloglevel

import (
	"flag"
)

var (
	verbose bool
	veryVerbose bool
	veryVeryVerbose bool
	veryVeryVeryVerbose bool
	veryVeryVeryVeryVerbose bool
	veryVeryVeryVeryVeryVerbose bool
)

func init() {

	flag.BoolVar(&verbose,                          "v", false, "verbose")
	flag.BoolVar(&veryVerbose,                     "vv", false, "very verbose")
	flag.BoolVar(&veryVeryVerbose,                "vvv", false, "very very verbose")
	flag.BoolVar(&veryVeryVeryVerbose,           "vvvv", false, "very very very verbose")
	flag.BoolVar(&veryVeryVeryVeryVerbose,      "vvvvv", false, "very very very very verbose")
	flag.BoolVar(&veryVeryVeryVeryVeryVerbose, "vvvvvv", false, "very very very very very verbose")
}

func Receive(dst *uint8) error {
	if !flag.Parsed() {
		return errNotParsedYet
	}

	if nil == dst {
		return errNilDestination
	}

	var value uint8
	{
		// -v -vv -vvv -vvvv -vvvvv -vvvvvv
		switch {
		case veryVeryVeryVeryVeryVerbose:
			value = 6
		case veryVeryVeryVeryVerbose:
			value = 5
		case veryVeryVeryVerbose:
			value = 4
		case veryVeryVerbose:
			value = 3
		case veryVerbose:
			value = 2
		case verbose:
			value = 1
		default:
			value = 0
		}
	}

        *dst = value
	return nil
}
