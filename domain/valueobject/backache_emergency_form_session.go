package valueobject

var _ EmergencyFormSession = (*BackacheEmergencyFormSession)(nil)

type BackacheEmergencyFormSession struct {
	Has *bool
}

func (f BackacheEmergencyFormSession) IsSet() bool { return f.Has != nil }

func (f BackacheEmergencyFormSession) Score() float64 {
	if !f.IsSet() || !*f.Has {
		return 0
	}

	return 1
}
