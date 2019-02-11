package projectcfg

type DeployHostName string

type DeployHostMap map[HostName]*DeployHost

type DeployHosts []*Host

type DeployHost struct {
	Branch BranchName `json:"branch,omitempty"`
	Before string     `json:"before,omitempty"`
	After  string     `json:"after,omitempty"`
	Files  FileMap    `json:"files,omitempty"`
}
