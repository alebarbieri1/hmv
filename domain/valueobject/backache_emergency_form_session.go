package valueobject

var _ EmergencyFormSession = (*BackacheEmergencyFormSession)(nil)

// BackacheEmergencyFormSession is an EmergencyForm session restricted for backache information
type BackacheEmergencyFormSession struct {
	Has *bool
}

// IsSet returns true if EmergencyFormSession is set
func (f BackacheEmergencyFormSession) IsSet() bool { return f.Has != nil }

// Score returns a float64 score according to the EmergencyFormSession level of criticity
func (f BackacheEmergencyFormSession) Score() float64 {
	if f.IsSet() && *f.Has {
		return 1
	}

	return 0
}
