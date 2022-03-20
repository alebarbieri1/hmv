package valueobject

var _ EmergencyFormSession = (*BreathingEmergencyFormSession)(nil)

type BreathingEmergencyFormSession struct {
	hasBreathingDifficulties *bool
}

func NewBreathingEmergencyFormSession(hasBreathingDifficulties bool) BreathingEmergencyFormSession {
	return BreathingEmergencyFormSession{hasBreathingDifficulties: &hasBreathingDifficulties}
}

func (f BreathingEmergencyFormSession) IsSet() bool { return f.hasBreathingDifficulties != nil }

func (f BreathingEmergencyFormSession) Score() float64 {
	if !f.IsSet() || !*f.hasBreathingDifficulties {
		return 0
	}

	return 1
}
