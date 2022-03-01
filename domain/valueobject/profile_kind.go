package valueobject

type ProfileKind string

func (p ProfileKind) String() string { return string(p) }

const (
	Analyst_ProfileKind   ProfileKind = "analyst"
	Pacient_ProfileKind   ProfileKind = "pacient"
	Undefined_ProfileKind ProfileKind = "undefined"
)
