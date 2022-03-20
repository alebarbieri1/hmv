package valueobject

var _ EmergencyFormSession = (*BreathingDifficultiesEmergencyFormSession)(nil)

type BreathingDifficultiesEmergencyFormSession struct {
	Has *bool
}

func (f BreathingDifficultiesEmergencyFormSession) IsSet() bool {
	return f.Has != nil
}

func (f BreathingDifficultiesEmergencyFormSession) Score() float64 {
	if !f.IsSet() || !*f.Has {
		return 0
	}

	return 1
}
