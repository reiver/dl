package argtarget

import (
	"github.com/reiver/dl/lib/opt/str"

	"flag"
)

var (
	value optstr.String
)

func init() {
	flag.Var(&value, "target", "The target to download — this could just be a URI")
}


func Receive(dst *optstr.String) error {

	if !flag.Parsed() {
		return errNotParsedYet
	}

	if nil == dst {
		return errNilDestination
	}

	var src optstr.String = value

	if optstr.Nothing() != src {
		*dst = src
		return nil
	}

	func(){
		if 0 >= flag.NArg() {
			return
		}

		src = optstr.Something(flag.Arg(0))
	}()

	return errNotFound
}
