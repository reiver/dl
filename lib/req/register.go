package req

func Register(scheme string, defaultPort int, requestor Requestor) error {
	return DefaultRegistry.Register(scheme, defaultPort, requestor)
}
