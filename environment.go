package projectcfg

type EnvironmentName string
type EnvironmentMap map[EnvironmentName]*Environment
type Environments []*Environment
type Environment struct {
	Name     string   `json:"name,omitempty"`
	Version  string   `json:"version,omitempty"`
	Hostname HostName `json:"hostname,omitempty"`
	Aliases  Aliases  `json:"aliases,omitempty"`
}
