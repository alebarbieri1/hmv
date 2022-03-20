package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*ChestPainEmergencyFormSession)(nil)

type ChestPainEmergencyFormSession struct {
	Has             *bool
	Characteristics ChestPainCharacteristics
}

func (f ChestPainEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.Has),
		ozzo.Field(&f.Characteristics, ozzo.By(func(value interface{}) error {
			if !f.IsSet() || !*f.Has {
				return ozzo.In(Undefined_ChestPainCharacteristics).Validate(value)
			}

			return ozzo.NotIn(Undefined_ChestPainCharacteristics).Validate(value)
		})),
	)
}

func (f ChestPainEmergencyFormSession) IsSet() bool { return f.Has != nil }

func (f ChestPainEmergencyFormSession) Score() float64 {
	if !f.IsSet() || !*f.Has {
		return Undefined_ChestPainCharacteristics.Float64()
	}

	return f.Characteristics.Float64() / RadiatingToTheLeftArm_ChestPainCharacteristics.Float64()
}
