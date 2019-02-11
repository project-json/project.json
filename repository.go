package projectcfg

type Repositories []*Repository

type Repository struct {
	Provider string   `cfg:"provider,omitempty"`
	Url      string   `cfg:"url"`
	Mirrors  []string `cfg:"mirrors,omitempty"`
}
