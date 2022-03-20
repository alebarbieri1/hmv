package valueobject

var _ EmergencyFormSession = (*BackacheEmergencyFormSession)(nil)

type BackacheEmergencyFormSession struct {
	hasBackache *bool
}

func NewBackacheEmergencyFormSession(hasBackache bool) BackacheEmergencyFormSession {
	return BackacheEmergencyFormSession{hasBackache: &hasBackache}
}

func (f BackacheEmergencyFormSession) IsSet() bool { return f.hasBackache != nil }

func (f BackacheEmergencyFormSession) Score() float64 {
	if !f.IsSet() || !*f.hasBackache {
		return 0
	}

	return 1
}
