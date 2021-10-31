package req

import (
	"io"
)

type RequestorFunc func(string, string) (io.ReadCloser, error)

func (fn RequestorFunc) Request(addr string, target string) (io.ReadCloser, error) {
	return fn(addr, target)
}
