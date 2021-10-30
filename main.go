package main

import (
	"github.com/reiver/dl/arg"
	"github.com/reiver/dl/lib/help"
	"github.com/reiver/dl/lib/opt/str"

	"fmt"
	"os"
)

func main() {

	// If the user has told us to output log messages (that can, for example be used for debugging) then
	// output the (command line) arguments we were given (as key-value pairs).
	if 0 < arg.LogLevel {
		fmt.Fprintf(os.Stderr, "help      🡆 %#t\n", arg.Help)
		fmt.Fprintf(os.Stderr, "host      🡆 %#v\n", arg.Host)
		fmt.Fprintf(os.Stderr, "log-level 🡆 %d\n", arg.LogLevel)
		fmt.Fprintf(os.Stderr, "scheme    🡆 %#v\n", arg.Scheme)
		fmt.Fprintf(os.Stderr, "target    🡆 %#v\n", arg.Target)
		fmt.Fprintf(os.Stderr, "version   🡆 %t\n", arg.Version)
	}

	// If the user asked for the help message to be outputted (by calling this program with the --help flag) then
	// output the help message (on STDERR), and exit with an exit code of ‘OK’ — i.e., 0 (zero).
	if arg.Help {
		const exitCodeOK = 0

		help.WriteTo(os.Stderr)
		os.Exit(exitCodeOK)
		return
	}

	// If the user didn't provide any flags, switches, or parameters, on the command line then
	// output the help message (on STDERR), and exit with an exit code of ‘usage-error’ — i.e., 64 (sixty-four).
	if optstr.Nothing() == arg.Target {
		const exitCodeUsageError = 64

		help.WriteTo(os.Stderr)
		os.Exit(exitCodeUsageError)
		return
	}

	// If the user provided the ‘target’ then we can (elsewhere in the code) infer the ‘scheme’ and ‘host’.
	// Also, the user can override the ‘scheme’ and ‘host’ (inferred from the ‘target’).
	//
	// Thus, if for some weird reason we have the ‘target’ but don't have the ‘scheme’ or ‘host’ then
	// some weird internal error (like a bug) happened so exit with an exit code of ‘internal-software-error’ — i.e., 70 (seventy).
	if optstr.Nothing() == arg.Scheme {
		const exitCodeInternalSoftwareError = 70

		fmt.Println(os.Stderr, "uh oh! — internal software error: could not figure out the ‘scheme’")
		os.Exit(exitCodeInternalSoftwareError)
		return
	}
	if optstr.Nothing() == arg.Host {
		const exitCodeInternalSoftwareError = 70

		fmt.Println(os.Stderr, "uh oh! — internal software error: could not figure out the ‘host’")
		os.Exit(exitCodeInternalSoftwareError)
		return
	}
}
