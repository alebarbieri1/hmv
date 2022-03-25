package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*ChestPainEmergencyFormSession)(nil)

// ChestPainEmergencyFormSession is an EmergencyForm session restricted for chest pain information
type ChestPainEmergencyFormSession struct {
	Has             *bool
	Characteristics ChestPainCharacteristics
}

// Validate validates a ChestPainEmergencyFormSession
func (f ChestPainEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.Has),
		ozzo.Field(&f.Characteristics, ozzo.By(func(value interface{}) error {
			if f.IsSet() && *f.Has {
				return ozzo.Required.Validate(value)
			}

			return ozzo.In(Undefined_ChestPainCharacteristics).Validate(value)
		})),
	)
}

// IsSet returns true if EmergencyFormSession is set
func (f ChestPainEmergencyFormSession) IsSet() bool { return f.Has != nil }

// Score returns a float64 score according to the EmergencyFormSession level of criticity
func (f ChestPainEmergencyFormSession) Score() float64 {
	return f.Characteristics.Float64() / RadiatingToTheLeftArm_ChestPainCharacteristics.Float64()
}
