package projectcfg

import (
	"github.com/projectcfg/projectcfg/util/cardinality"
)

type StackService struct {
	cardinality.Cardinality
	*Service
}
