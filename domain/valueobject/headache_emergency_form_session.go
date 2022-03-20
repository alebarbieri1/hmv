package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*HeadacheEmergencyFormSession)(nil)

type HeadacheEmergencyFormSession struct {
	hasHeadache       *bool
	headacheIntensity HeadacheIntensity
}

func NewHeadacheEmergencyFormSession(hasHeadache bool, intensity HeadacheIntensity) HeadacheEmergencyFormSession {
	return HeadacheEmergencyFormSession{hasHeadache: &hasHeadache, headacheIntensity: intensity}
}

func (f HeadacheEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.hasHeadache, ozzo.By(func(value interface{}) error {
			if f.headacheIntensity == Undefined_HeadacheIntensity {
				return ozzo.Nil.Validate(value)
			}

			return ozzo.NotNil.Validate(value)
		})),
		ozzo.Field(&f.headacheIntensity, ozzo.By(func(value interface{}) error {
			if f.hasHeadache == nil || !*f.hasHeadache {
				return ozzo.In(Undefined_HeadacheIntensity).Validate(value)
			}

			return ozzo.NotIn(Undefined_HeadacheIntensity).Validate(value)
		})),
	)
}

func (f HeadacheEmergencyFormSession) IsSet() bool { return f.hasHeadache != nil }

func (f HeadacheEmergencyFormSession) Score() float64 {
	if !f.IsSet() {
		return Undefined_HeadacheIntensity.Float64()
	}

	return f.headacheIntensity.Float64() / VeryHigh_HeadacheIntensity.Float64()
}
