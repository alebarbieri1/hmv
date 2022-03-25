package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*HeadacheEmergencyFormSession)(nil)

// HeadacheEmergencyFormSession is an EmergencyForm session restricted for headache information
type HeadacheEmergencyFormSession struct {
	Has       *bool
	Intensity HeadacheIntensity
}

// Validate validates a HeadacheEmergencyFormSession
func (f HeadacheEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.Has),
		ozzo.Field(&f.Intensity, ozzo.By(func(value interface{}) error {
			if f.IsSet() && *f.Has {
				return ozzo.Required.Validate(value)
			}

			return ozzo.In(Undefined_HeadacheIntensity).Validate(value)
		})),
	)
}

// IsSet returns true if EmergencyFormSession is set
func (f HeadacheEmergencyFormSession) IsSet() bool { return f.Has != nil }

// Score returns a float64 score according to the EmergencyFormSession level of criticity
func (f HeadacheEmergencyFormSession) Score() float64 {
	return f.Intensity.Float64() / VeryHigh_HeadacheIntensity.Float64()
}
