package projectcfg

import (
	"github.com/projectcfg/projectcfg/util/cardinality"
)

type StackRoleIdentifier string

type StackRole struct {
	NamedStack string
	*Role
	Cardinality cardinality.Cardinality
}
