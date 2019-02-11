package projectcfg

type Repositories []*Repository

type Repository struct {
	Provider string   `json:"provider,omitempty"`
	Url      string   `json:"url"`
	Mirrors  []string `json:"mirrors,omitempty"`
}
