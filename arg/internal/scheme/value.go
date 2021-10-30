package argscheme

import (
	"github.com/reiver/dl/arg/internal/target"
	"github.com/reiver/dl/lib/opt/str"

	"flag"
	"net/url"
)

var (
	value optstr.String
)

func init() {
	flag.Var(&value, "scheme", "The scheme to use  — ex: gemini, gopher, http, https, mercury, scp, smb")
}


func Receive(dst *optstr.String) error {

	if !flag.Parsed() {
		return errNotParsedYet
	}

	if nil == dst {
		return errNilDestination
	}

	if optstr.Nothing() != value {
		*dst = value
		return nil
	}

	var src optstr.String
	func(){
		var target optstr.String
		{
			err := argtarget.Receive(&target)
			if nil != err {
				panic(err)
			}
		}

		if optstr.Nothing() == target {
			return
		}

		uri, err := url.Parse(target.String())
		if nil == err {
			x := uri.Scheme
			if "" != x {
				src = optstr.Something(x)
			}
		}
	}()

	*dst = src
	return nil
}
