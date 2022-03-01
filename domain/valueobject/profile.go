package valueobject

type Profile string

func (p Profile) String() string { return string(p) }

const (
	AnalystProfile Profile = "analyst"
	PacientProfile Profile = "pacient"
)
