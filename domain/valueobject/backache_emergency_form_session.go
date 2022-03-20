package valueobject

var _ EmergencyFormSession = (*BackacheEmergencyFormSession)(nil)

type BackacheEmergencyFormSession struct {
	HasBackache *bool
}

func (f BackacheEmergencyFormSession) IsSet() bool { return f.HasBackache != nil }

func (f BackacheEmergencyFormSession) Score() float64 {
	if !f.IsSet() || !*f.HasBackache {
		return 0
	}

	return 1
}
