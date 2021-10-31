package verboten

import (
	"github.com/reiver/dl/lib/req"

	"github.com/reiver/go-hg"

	"io"
)

const scheme      = "mercury"
const defaultPort = 1961

func init() {

	var requestor req.Requestor = req.RequestorFunc(dorequest)

	err := req.Register(scheme, defaultPort, requestor)
	if nil != err {
		panic(err)
	}
}


func dorequest(addr string, target string) (io.ReadCloser, error) {

	var request hg.Request
	{
		err := request.Parse(target)
		if nil != err {
			return nil, err
		}
	}

	var rr hg.ResponseReader
	{
		var err error

		rr, err = hg.DialAndCall(addr, request)
		if nil != err {
			return nil, err
		}

		if nil == rr {
			return nil, errNilResponseReader
		}
	}

	return rr, nil
}
