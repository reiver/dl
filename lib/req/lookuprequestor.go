package req

func LookupRequestor(scheme string) (Requestor, bool) {
	return DefaultRegistry.LookupRequestor(scheme)
}
