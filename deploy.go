package projectcfg

type Deploy struct {
	Provider    string        `cfg:"provider"`
	TargetName  string        `cfg:"target_name"`
	TargetId    string        `cfg:"target_id,omitempty"`
	WebRoot     string        `cfg:"web_root,omitempty"`
	Credentials Credentials   `cfg:"credentials,omitempty"`
	Frameworks  FrameworkMap  `cfg:"frameworks,omitempty"`
	Files       FileMap       `cfg:"files,omitempty"`
	Hosts       DeployHostMap `cfg:"hosts"`
}
