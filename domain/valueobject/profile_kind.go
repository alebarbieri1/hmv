package valueobject

// ProfileKind represents the user's profile kind
type ProfileKind string

// String returns the string value of an ProfileKind
func (p ProfileKind) String() string { return string(p) }

const (
	Analyst_ProfileKind   ProfileKind = "analyst"
	Pacient_ProfileKind   ProfileKind = "pacient"
	Rescuer_ProfileKind   ProfileKind = "rescuer"
	Undefined_ProfileKind ProfileKind = "undefined"
)
