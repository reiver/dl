package req

var DefaultRegistry Registry

func init() {
	DefaultRegistry = &internalRegistry{}
}
