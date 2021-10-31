package arg

import (
	"github.com/reiver/dl/lib/opt/str"

	"github.com/reiver/dl/arg/internal/address"
	"github.com/reiver/dl/arg/internal/help"
	"github.com/reiver/dl/arg/internal/loglevel"
	"github.com/reiver/dl/arg/internal/scheme"
	"github.com/reiver/dl/arg/internal/target"
	"github.com/reiver/dl/arg/internal/version"

	"flag"
)

var (
	Address optstr.String
	Help bool
	LogLevel uint8 = 0
	Scheme optstr.String
	Target optstr.String
	Version bool
)

func init() {

	flag.Parse()

	{
		var err error

		// -- help
		err = arghelp.Receive(&Help)
		if nil != err {
			panic(err)
		}

		// -v -vv -vvv -vvvv -vvvvv -vvvvvv
		err = argloglevel.Receive(&LogLevel)
		if nil != err {
			panic(err)
		}

		// --address
		err = argaddress.Receive(&Address)
		if nil != err {
			panic(err)
		}

		// --scheme
		err = argscheme.Receive(&Scheme)
		if nil != err {
			panic(err)
		}

		// --target
		err = argtarget.Receive(&Target)
		if nil != err {
			panic(err)
		}

		// --version
		err = argversion.Receive(&Version)
		if nil != err {
			panic(err)
		}
	}
}
