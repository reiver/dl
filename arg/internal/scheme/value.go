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

	var fromtarget optstr.String
	{
		err := func() error {
			var target optstr.String
			{
				err := argtarget.Receive(&target)
				if nil != err {
					return err
				}
			}

			if optstr.Nothing() == target {
				return nil
			}

			uri, err := url.Parse(target.String())
			if nil == err {
				x := uri.Scheme
				if "" != x {
					fromtarget = optstr.Something(x)
					return nil
				}
			}

			return nil
		}()

		if nil != err {
			return err
		}
	}

	optstr.Push(dst, value, fromtarget)

	return nil
}
