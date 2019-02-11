package projectcfg

type Source struct {
	WebRoot      string       `cfg:"web_root,omitempty"`
	Frameworks   FrameworkMap `cfg:"frameworks,omitempty"`
	Repositories Repositories `cfg:"repositories,omitempty"`
	Branches     BranchMap    `cfg:"branches,omitempty"`
}
