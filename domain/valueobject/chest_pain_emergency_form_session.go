package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*ChestPainEmergencyFormSession)(nil)

type ChestPainEmergencyFormSession struct {
	HasChestPain             *bool
	ChestPainCharacteristics ChestPainCharacteristics
}

func (f ChestPainEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.HasChestPain, ozzo.By(func(value interface{}) error {
			if f.ChestPainCharacteristics == Undefined_ChestPainCharacteristics {
				return ozzo.Nil.Validate(value)
			}

			return ozzo.NotNil.Validate(value)
		})),
		ozzo.Field(&f.ChestPainCharacteristics, ozzo.By(func(value interface{}) error {
			if !f.IsSet() || !*f.HasChestPain {
				return ozzo.In(Undefined_ChestPainCharacteristics).Validate(value)
			}

			return ozzo.NotIn(Undefined_ChestPainCharacteristics).Validate(value)
		})),
	)
}

func (f ChestPainEmergencyFormSession) IsSet() bool { return f.HasChestPain != nil }

func (f ChestPainEmergencyFormSession) Score() float64 {
	if !f.IsSet() || !*f.HasChestPain {
		return Undefined_ChestPainCharacteristics.Float64()
	}

	return f.ChestPainCharacteristics.Float64() / RadiatingToTheLeftArm_ChestPainCharacteristics.Float64()
}
