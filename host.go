package projectcfg

type HostName string

type HostMap map[HostName]*Host

type Hosts []*Host

type Host struct {
	Provider    string       `cfg:"provider,omitempty"`
	Name        string       `cfg:"name,omitempty"`
	Label       string       `cfg:"label,omitempty"`
	WebRoot     string       `cfg:"web_root,omitempty"`
	Credentials *Credentials `cfg:"credentials,omitempty"`
	Repository  *Repository  `cfg:"repository,omitempty"`
}
