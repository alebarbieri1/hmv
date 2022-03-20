package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*OxygenSaturationEmergencyFormSession)(nil)

// OxygenSaturationEmergencyFormSession is a form session data about oxygen saturation
// - OxygenSaturation: 0 <= value <= 1
type OxygenSaturationEmergencyFormSession struct {
	Value *float64
}

func (f OxygenSaturationEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.Value, ozzo.Min(0), ozzo.Max(1)),
	)
}

func (f OxygenSaturationEmergencyFormSession) IsSet() bool { return f.Value != nil }

func (f OxygenSaturationEmergencyFormSession) Score() float64 {
	saturarion := Undefined_OxygenSaturationLevel

	if !f.IsSet() {
		return saturarion.Float64()
	}

	switch {
	case *f.Value < 0.85:
		saturarion = SeverelyHypoxic_OxygenSaturationLevel
	case 0.85 <= *f.Value && *f.Value < 0.95:
		saturarion = Hypoxic_OxygenSaturationLevel
	case *f.Value >= 0.95:
		saturarion = Normal_OxygenSaturationLevel
	}

	return saturarion.Float64() / SeverelyHypoxic_OxygenSaturationLevel.Float64()
}
