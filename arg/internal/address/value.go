package argaddress

import (
	"github.com/reiver/dl/arg/internal/target"
	"github.com/reiver/dl/lib/opt/str"
	"github.com/reiver/dl/lib/req"

	"flag"
	"fmt"
	"net"
	"net/url"
)

var (
	value optstr.String
)

func init() {
	flag.Var(&value, "address", "The address (host & port) to make the request from  — ex: “example.com:1961”, “192.0.2.1:8008”")
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
				host := uri.Hostname()
				if "" == host {
					return nil
				}

				port := uri.Port()
				if "" == port {
					port = req.LookupDefaultPort(uri.Scheme).String()
				}
				if "" == port {
					portint, err := net.LookupPort("tcp", uri.Scheme)
					if nil == err {
						port = fmt.Sprint(portint)
					}
				}
				if "" == port {
					return nil
				}

				var addr string = net.JoinHostPort(host, port)

				fromtarget = optstr.Something(addr)

				return nil
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
