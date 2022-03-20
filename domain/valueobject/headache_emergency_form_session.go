package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*HeadacheEmergencyFormSession)(nil)

type HeadacheEmergencyFormSession struct {
	HasHeadache       *bool
	HeadacheIntensity HeadacheIntensity
}

func (f HeadacheEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.HasHeadache, ozzo.By(func(value interface{}) error {
			if f.HeadacheIntensity == Undefined_HeadacheIntensity {
				return ozzo.Nil.Validate(value)
			}

			return ozzo.NotNil.Validate(value)
		})),
		ozzo.Field(&f.HeadacheIntensity, ozzo.By(func(value interface{}) error {
			if f.HasHeadache == nil || !*f.HasHeadache {
				return ozzo.In(Undefined_HeadacheIntensity).Validate(value)
			}

			return ozzo.NotIn(Undefined_HeadacheIntensity).Validate(value)
		})),
	)
}

func (f HeadacheEmergencyFormSession) IsSet() bool { return f.HasHeadache != nil }

func (f HeadacheEmergencyFormSession) Score() float64 {
	if !f.IsSet() {
		return Undefined_HeadacheIntensity.Float64()
	}

	return f.HeadacheIntensity.Float64() / VeryHigh_HeadacheIntensity.Float64()
}
