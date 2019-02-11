package projectcfg

type BranchName string

type BranchMap map[BranchName]*Branch
type Branches []*Branch

type Branch struct {
	Name BranchName `cfg:"branch,omitempty"`
}
