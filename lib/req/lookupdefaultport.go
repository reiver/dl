package req

import (
	"github.com/reiver/dl/lib/opt/int"
)

func LookupDefaultPort(scheme string) optint.Int {
	return DefaultRegistry.LookupDefaultPort(scheme)
}
