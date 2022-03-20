package valueobject

var _ EmergencyFormSession = (*BreathingEmergencyFormSession)(nil)

type BreathingEmergencyFormSession struct {
	HasBreathingDifficulties *bool
}

func (f BreathingEmergencyFormSession) IsSet() bool { return f.HasBreathingDifficulties != nil }

func (f BreathingEmergencyFormSession) Score() float64 {
	if !f.IsSet() || !*f.HasBreathingDifficulties {
		return 0
	}

	return 1
}
