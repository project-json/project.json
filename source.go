package projectcfg

type Source struct {
	WebRoot      string       `json:"web_root,omitempty"`
	Frameworks   FrameworkMap `json:"frameworks,omitempty"`
	Repositories Repositories `json:"repositories,omitempty"`
	Branches     BranchMap    `json:"branches,omitempty"`
}
