package valueobject

import (
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

var _ EmergencyFormSession = (*AbdominalPainEmergencyFormSession)(nil)

// AbdominalPainEmergencyFormSession is an EmergencyForm session restricted for abdominal pain information
type AbdominalPainEmergencyFormSession struct {
	Has       *bool
	Intensity AbdominalPainIntensity
}

// Validate validates an AbdominalPainEmergencyFormSession
func (f AbdominalPainEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.Has),
		ozzo.Field(&f.Intensity, ozzo.By(func(value interface{}) error {
			if f.IsSet() && *f.Has {
				return ozzo.Required.Validate(value)
			}

			return ozzo.In(Undefined_AbdominalPainIntensity).Validate(value)
		})),
	)
}

// IsSet returns true if EmergencyFormSession is set
func (f AbdominalPainEmergencyFormSession) IsSet() bool { return f.Has != nil }

// Score returns a float64 score according to the EmergencyFormSession level of criticity
func (f AbdominalPainEmergencyFormSession) Score() float64 {
	return f.Intensity.Float64() / VeryHigh_AbdominalPainIntensity.Float64()
}
