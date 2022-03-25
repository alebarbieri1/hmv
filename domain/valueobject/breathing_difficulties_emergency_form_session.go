package valueobject

var _ EmergencyFormSession = (*BreathingDifficultiesEmergencyFormSession)(nil)

// BreathingDifficultiesEmergencyFormSession is an EmergencyForm session restricted for breathing difficulties information
type BreathingDifficultiesEmergencyFormSession struct {
	Has *bool
}

// IsSet returns true if EmergencyFormSession is set
func (f BreathingDifficultiesEmergencyFormSession) IsSet() bool { return f.Has != nil }

// Score returns a float64 score according to the EmergencyFormSession level of criticity
func (f BreathingDifficultiesEmergencyFormSession) Score() float64 {
	if f.IsSet() && *f.Has {
		return 1
	}

	return 0
}
