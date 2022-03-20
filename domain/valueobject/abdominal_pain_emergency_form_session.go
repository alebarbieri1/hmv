package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*AbdominalPainEmergencyFormSession)(nil)

type AbdominalPainEmergencyFormSession struct {
	Has       *bool
	Intensity AbdominalPainIntensity
}

func (f AbdominalPainEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.Has),
		ozzo.Field(&f.Intensity, ozzo.By(func(value interface{}) error {
			if !f.IsSet() || !*f.Has {
				return ozzo.In(Undefined_AbdominalPainIntensity).Validate(value)
			}

			return ozzo.NotIn(Undefined_AbdominalPainIntensity).Validate(value)
		})),
	)
}

func (f AbdominalPainEmergencyFormSession) IsSet() bool { return f.Has != nil }

func (f AbdominalPainEmergencyFormSession) Score() float64 {
	if !f.IsSet() {
		return 0
	}

	return f.Intensity.Float64() / VeryHigh_AbdominalPainIntensity.Float64()
}
