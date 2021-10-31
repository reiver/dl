package req

import (
	"errors"
)

var (
	errFound        = errors.New("req: found")
	errNilReceiver  = errors.New("req: nil receiver")
	errNilRequestor = errors.New("req: nil requestor")
)
