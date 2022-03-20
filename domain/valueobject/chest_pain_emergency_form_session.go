package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*ChestPainEmergencyFormSession)(nil)

type ChestPainEmergencyFormSession struct {
	hasChestPain             *bool
	chestPainCharacteristics ChestPainCharacteristics
}

func NewChestPainEmergencyFormSession(hasChestPain bool, chestPainCharacteristics ChestPainCharacteristics) ChestPainEmergencyFormSession {
	return ChestPainEmergencyFormSession{hasChestPain: &hasChestPain, chestPainCharacteristics: chestPainCharacteristics}
}

func (f ChestPainEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.hasChestPain, ozzo.By(func(value interface{}) error {
			if f.chestPainCharacteristics == Undefined_ChestPainCharacteristics {
				return ozzo.Nil.Validate(value)
			}

			return ozzo.NotNil.Validate(value)
		})),
		ozzo.Field(&f.chestPainCharacteristics, ozzo.By(func(value interface{}) error {
			if !f.IsSet() || !*f.hasChestPain {
				return ozzo.In(Undefined_ChestPainCharacteristics).Validate(value)
			}

			return ozzo.NotIn(Undefined_ChestPainCharacteristics).Validate(value)
		})),
	)
}

func (f ChestPainEmergencyFormSession) IsSet() bool { return f.hasChestPain != nil }

func (f ChestPainEmergencyFormSession) Score() float64 {
	if !f.IsSet() || !*f.hasChestPain {
		return Undefined_ChestPainCharacteristics.Float64()
	}

	return f.chestPainCharacteristics.Float64() / RadiatingToTheLeftArm_ChestPainCharacteristics.Float64()
}
