package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*OxygenSaturationEmergencyFormSession)(nil)

// OxygenSaturationEmergencyFormSession is a form session data about oxygen saturation
// - OxygenSaturation: 0 <= value <= 1
type OxygenSaturationEmergencyFormSession struct {
	Value *float64
}

// Validate validates a OxygenSaturationEmergencyFormSession
func (f OxygenSaturationEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.Value, ozzo.Min(0.0), ozzo.Max(1.0)),
	)
}

// IsSet returns true if EmergencyFormSession is set
func (f OxygenSaturationEmergencyFormSession) IsSet() bool { return f.Value != nil }

// Score returns a float64 score according to the EmergencyFormSession level of criticity
func (f OxygenSaturationEmergencyFormSession) Score() float64 {
	if f.IsSet() {
		switch {
		case *f.Value < 0.85:
			return 1.0
		case 0.85 <= *f.Value && *f.Value < 0.95:
			return 0.75
		}
	}

	return 0
}
