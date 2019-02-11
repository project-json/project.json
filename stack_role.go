package projectcfg

type StackRoleIdentifier string

type StackRole struct {
	NamedStack string
	*Role
	Ordinality Ordinality
}
