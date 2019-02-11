package projectcfg

type ProjectType string

const (
	WebsiteProject ProjectType = "website"
)

type Ordinality int

const (
	UnknownOrdinality Ordinality = iota
	OnlyOne
	OneOrMore
	ZeroOrMore
	TwoOrMore
)
