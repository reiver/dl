package optstr

// Push iterates through ‘sources’, and the first non-Nothing source will get pushed into ‘destination’.
//
// So, for example, if this was called with:
//
//	succeeded := optstr.Push(&dst,
//		optstr.Nothing(),
//		optstr.Nothing(),
//		optstr.Something("once"),
//		optstr.Nothing(),
//		optstr.Nothing(),
//		optstr.Something("twice"),
//		optstr.Nothing(),
//	)
//
// Then optstr.Something("once") would get pushed into ‘destination’.
func Push(destination *String, sources ...String) bool {

	if nil == destination {
		return false
	}

	for _, source := range sources {
		if Nothing() != source {
			*destination = source
			return true
		}
	}

	return false
}
