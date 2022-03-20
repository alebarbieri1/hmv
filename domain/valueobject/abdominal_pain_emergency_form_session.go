package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*AbdominalPainEmergencyFormSession)(nil)

type AbdominalPainEmergencyFormSession struct {
	hasAbdominalPain       *bool
	abdominalPainIntensity AbdominalPainIntensity
}

func NewAbdominalPainEmergencyFormSession(hasAbdominalPain bool, abdominalPainIntensity AbdominalPainIntensity) AbdominalPainEmergencyFormSession {
	return AbdominalPainEmergencyFormSession{hasAbdominalPain: &hasAbdominalPain, abdominalPainIntensity: abdominalPainIntensity}
}

func (f AbdominalPainEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.hasAbdominalPain, ozzo.By(func(value interface{}) error {
			if f.abdominalPainIntensity == Undefined_AbdominalPainIntensity {
				return ozzo.Nil.Validate(value)
			}

			return ozzo.NotNil.Validate(value)
		})),
		ozzo.Field(&f.abdominalPainIntensity, ozzo.By(func(value interface{}) error {
			if !f.IsSet() || !*f.hasAbdominalPain {
				return ozzo.In(Undefined_AbdominalPainIntensity).Validate(value)
			}

			return ozzo.NotIn(Undefined_AbdominalPainIntensity).Validate(value)
		})),
	)
}

func (f AbdominalPainEmergencyFormSession) IsSet() bool { return f.hasAbdominalPain != nil }

func (f AbdominalPainEmergencyFormSession) Score() float64 {
	if !f.IsSet() {
		return 0
	}

	return f.abdominalPainIntensity.Float64() / VeryHigh_AbdominalPainIntensity.Float64()
}
