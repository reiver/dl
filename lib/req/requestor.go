package req

import (
	"io"
)

type Requestor interface {
	Request(addr string, target string) (io.ReadCloser, error)
}
