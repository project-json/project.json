package projectcfg

type HostName string

type HostMap map[HostName]*Host

type Hosts []*Host

type Host struct {
	Provider    string       `json:"provider,omitempty"`
	Name        string       `json:"name,omitempty"`
	Label       string       `json:"label,omitempty"`
	WebRoot     string       `json:"web_root,omitempty"`
	Credentials *Credentials `json:"credentials,omitempty"`
	Repository  *Repository  `json:"repository,omitempty"`
}
