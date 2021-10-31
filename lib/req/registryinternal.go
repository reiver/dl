package req

import (
	"github.com/reiver/dl/lib/opt/int"

	"sync"
)

type internalRegistry struct {
	mutex sync.Mutex
	requestors map[string]Requestor
	ports map[string]int
}

func (receiver *internalRegistry) Register(scheme string, port int, requestor Requestor) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil == requestor {
		return  errNilRequestor
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.requestors {
		receiver.requestors = map[string]Requestor{}
	}
	if nil == receiver.ports {
		receiver.ports = map[string]int{}
	}

	{
		_, found := receiver.requestors[scheme]
		if found {
			return errFound
		}
	}
	{
		_, found := receiver.ports[scheme]
		if found {
			return errFound
		}
	}

	receiver.requestors[scheme] = requestor
	receiver.ports[scheme]      = port

	return nil
}

func (receiver *internalRegistry) LookupRequestor(scheme string) (Requestor, bool) {
	if nil == receiver {
		return nil, false
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	requestors := receiver.requestors

	if nil == requestors {
		return nil, false
	}

	value, found := requestors[scheme]
	if !found {
		return nil, false
	}

	return value, true
}

func (receiver *internalRegistry) LookupDefaultPort(scheme string) optint.Int {
	if nil == receiver {
		return optint.Nothing()
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	ports := receiver.ports

	if nil == ports {
		return optint.Nothing()
	}

	value, found := ports[scheme]
	if !found {
		return optint.Nothing()
	}

	return optint.Something(value)
}
