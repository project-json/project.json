package projectcfg

const DefaultFilename string = "project.json"

type ProjectType string

const (
	WebsiteProject ProjectType = "website"
	DocsProject    ProjectType = "docs"
	LibaryProject  ProjectType = "library"
)

type Ordinality int

const (
	UnknownOrdinality Ordinality = iota
	OnlyOne
	OneOrMore
	ZeroOrMore
	TwoOrMore
)
