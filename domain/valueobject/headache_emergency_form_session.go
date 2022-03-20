package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*HeadacheEmergencyFormSession)(nil)

type HeadacheEmergencyFormSession struct {
	Has       *bool
	Intensity HeadacheIntensity
}

func (f HeadacheEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.Has),
		ozzo.Field(&f.Intensity, ozzo.By(func(value interface{}) error {
			if f.Has == nil || !*f.Has {
				return ozzo.In(Undefined_HeadacheIntensity).Validate(value)
			}

			return ozzo.NotIn(Undefined_HeadacheIntensity).Validate(value)
		})),
	)
}

func (f HeadacheEmergencyFormSession) IsSet() bool { return f.Has != nil }

func (f HeadacheEmergencyFormSession) Score() float64 {
	if !f.IsSet() {
		return Undefined_HeadacheIntensity.Float64()
	}

	return f.Intensity.Float64() / VeryHigh_HeadacheIntensity.Float64()
}
