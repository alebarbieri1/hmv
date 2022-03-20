package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*AbdominalPainEmergencyFormSession)(nil)

type AbdominalPainEmergencyFormSession struct {
	HasAbdominalPain       *bool
	AbdominalPainIntensity AbdominalPainIntensity
}

func (f AbdominalPainEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.HasAbdominalPain, ozzo.By(func(value interface{}) error {
			if f.AbdominalPainIntensity == Undefined_AbdominalPainIntensity {
				return ozzo.Nil.Validate(value)
			}

			return ozzo.NotNil.Validate(value)
		})),
		ozzo.Field(&f.AbdominalPainIntensity, ozzo.By(func(value interface{}) error {
			if !f.IsSet() || !*f.HasAbdominalPain {
				return ozzo.In(Undefined_AbdominalPainIntensity).Validate(value)
			}

			return ozzo.NotIn(Undefined_AbdominalPainIntensity).Validate(value)
		})),
	)
}

func (f AbdominalPainEmergencyFormSession) IsSet() bool { return f.HasAbdominalPain != nil }

func (f AbdominalPainEmergencyFormSession) Score() float64 {
	if !f.IsSet() {
		return 0
	}

	return f.AbdominalPainIntensity.Float64() / VeryHigh_AbdominalPainIntensity.Float64()
}
