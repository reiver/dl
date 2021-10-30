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

	var arg0 optstr.String
	{
		if 1 <= flag.NArg() {
			arg0 = optstr.Something(flag.Arg(0))
		}
	}

	optstr.Push(dst, value, arg0)

	return nil
}
