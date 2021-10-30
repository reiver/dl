package main

import (
	"github.com/reiver/dl/arg"
	"github.com/reiver/dl/lib/help"
	"github.com/reiver/dl/lib/opt/str"

	"fmt"
	"os"
)

func main() {

	// If the user has told us to output log messages (that can, for example
	// be used for debugging) then output our arguments.

	if 0 < arg.LogLevel {
		fmt.Fprintf(os.Stderr, "help      🡆 %#t\n", arg.Help)
		fmt.Fprintf(os.Stderr, "host      🡆 %#v\n", arg.Host)
		fmt.Fprintf(os.Stderr, "log-level 🡆 %d\n", arg.LogLevel)
		fmt.Fprintf(os.Stderr, "scheme    🡆 %#v\n", arg.Scheme)
		fmt.Fprintf(os.Stderr, "target    🡆 %#v\n", arg.Target)
		fmt.Fprintf(os.Stderr, "version   🡆 %t\n", arg.Version)
	}

	if arg.Help {
		const exitCodeOK = 0

		help.WriteTo(os.Stderr)
		os.Exit(exitCodeOK)
		return
	}

	if optstr.Nothing() == arg.Target {
		const exitCodeUsageError = 64

		help.WriteTo(os.Stderr)
		os.Exit(exitCodeUsageError)
		return
	}
}
