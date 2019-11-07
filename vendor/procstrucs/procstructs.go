package procstrucs

type ArrayHosts []string
type ArrayVarnishUris []string
type ArrayPathUris []string

type OptionsPurge struct {
	HostGroup    []string `json:"hostgroup"`
	VarnishGroup []string `json:"varnishgroup"`
	PathGroup    []string `json:"pathgroup"`
}

func (i *ArrayHosts) String() string {
	return "Single Host"
}

func (i *ArrayHosts) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *ArrayVarnishUris) String() string {
	return "Single Varnish Uri"
}

func (i *ArrayVarnishUris) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *ArrayPathUris) String() string {
	return "Single Path Uri"
}

func (i *ArrayPathUris) Set(value string) error {
	*i = append(*i, value)
	return nil
}
