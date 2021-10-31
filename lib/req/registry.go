package req

import (
	"github.com/reiver/dl/lib/opt/int"
)

type Registry interface {
	Register(string, int, Requestor) error
	LookupRequestor(string) (Requestor, bool)
	LookupDefaultPort(string) optint.Int
}
