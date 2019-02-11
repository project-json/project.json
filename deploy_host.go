package projectcfg

type DeployHostName string

type DeployHostMap map[HostName]*DeployHost

type DeployHosts []*Host

type DeployHost struct {
	Branch BranchName `cfg:"branch,omitempty"`
	Before string     `cfg:"before,omitempty"`
	After  string     `cfg:"after,omitempty"`
	Files  FileMap    `cfg:"files,omitempty"`
}
