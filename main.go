package main

import (
	"github.com/reiver/dl/arg"
	"github.com/reiver/dl/lib/help"
	"github.com/reiver/dl/lib/opt/str"
	"github.com/reiver/dl/lib/req"

	"fmt"
	"io"
	"os"
)

func main() {

	// If the user has told us to output log messages (that can, for example be used for debugging) then
	// output the (command line) arguments we were given (as key-value pairs).
	if 0 < arg.LogLevel {
		fmt.Fprintf(os.Stderr, "address   🡆 %#v\n", arg.Address)
		fmt.Fprintf(os.Stderr, "help      🡆 %t\n",  arg.Help)
		fmt.Fprintf(os.Stderr, "log-level 🡆 %d\n",  arg.LogLevel)
		fmt.Fprintf(os.Stderr, "scheme    🡆 %#v\n", arg.Scheme)
		fmt.Fprintf(os.Stderr, "target    🡆 %#v\n", arg.Target)
		fmt.Fprintf(os.Stderr, "version   🡆 %t\n",  arg.Version)
	}

	// If the user asked for the help message to be outputted (by calling this program with the --help flag) then
	// output the help message (on STDERR), and exit with an exit-code of ‘OK’ — i.e., 0 (zero).
	if arg.Help {
		const exitCodeOK = 0

		help.WriteTo(os.Stderr)
		os.Exit(exitCodeOK)
		return
	}

	// If the user didn't provide any flags, switches, or parameters, on the command line then
	// output the help message (on STDERR), and exit with an exit-code of ‘usage-error’ — i.e., 64 (sixty-four).
	if optstr.Nothing() == arg.Target {
		const exitCodeUsageError = 64

		help.WriteTo(os.Stderr)
		os.Exit(exitCodeUsageError)
		return
	}

	// If the user provided the ‘target’ then we can (elsewhere in the code) infer the ‘scheme’ and ‘address’.
	// Also, the user can override the ‘scheme’ and ‘address’ (inferred from the ‘target’).
	//
	// Thus, if for some weird reason we have the ‘target’ but don't have the ‘scheme’ or ‘address’ then
	// some weird internal error (like a bug) happened so exit with an exit-code of ‘internal-software-error’ — i.e., 70 (seventy).
	if optstr.Nothing() == arg.Scheme {
		const exitCodeInternalSoftwareError = 70

		fmt.Fprintln(os.Stderr, "uh oh! — internal software error: could not figure out the ‘scheme’")
		os.Exit(exitCodeInternalSoftwareError)
		return
	}
	if optstr.Nothing() == arg.Address {
		const exitCodeInternalSoftwareError = 70

		fmt.Fprintln(os.Stderr, "uh oh! — internal software error: could not figure out the ‘address’")
		os.Exit(exitCodeInternalSoftwareError)
		return
	}

	var readcloser io.ReadCloser
	{
		requestor, found := req.LookupRequestor(arg.Scheme.String())
		if !found {
			fmt.Fprintf(os.Stderr, "sorry — that scheme (%s) is not supported — cannot request make a request to address %q for target %q\n", arg.Scheme, arg.Address, arg.Target)
			os.Exit(1)
			return
		}
		if nil == requestor {
			const exitCodeInternalSoftwareError = 70

			fmt.Fprintf(os.Stderr, "uh oh! — internal software error: was not given the requestor for %q\n", arg.Scheme)
			os.Exit(exitCodeInternalSoftwareError)
			return
		}

		var err error

		readcloser, err = requestor.Request(arg.Address.String(), arg.Target.String())
		if nil != err {
			const exitCodeServiceUnavailable = 69

			fmt.Fprintf(os.Stderr, "uh oh! — received an error when trying get data-connection to %q for %q: %s\n", arg.Address, arg.Target, err)
			os.Exit(exitCodeServiceUnavailable)
			return
		}
		defer func() {
			err := readcloser.Close()
			if nil != err {
				fmt.Fprintf(os.Stderr, "uh oh! — received an error when trying to close data-connection to %q for %q: %s\n", arg.Address, arg.Target, err)
			}
		}()
	}

	{
		_, err := io.Copy(os.Stdout, readcloser)
		if nil != err {
			const exitCodeTemporaryFailure = 75

			fmt.Fprintf(os.Stderr, "uh oh! — received an error when trying to send the data downloaded from the request to %q for %q to STDOUT: %s\n", arg.Address, arg.Target, err)
			os.Exit(exitCodeTemporaryFailure)
			return
		}
	}
}
