package projectcfg

type Environment struct {
	Name    string `cfg:"name,omitempty"`
	Version string `cfg:"version,omitempty"`
	Website string `cfg:"website,omitempty"`
}
