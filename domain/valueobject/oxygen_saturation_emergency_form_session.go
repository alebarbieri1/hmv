package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

var _ EmergencyFormSession = (*OxygenSaturationEmergencyFormSession)(nil)

// OxygenSaturationEmergencyFormSession is a form session data about oxygen saturation
// - oxygenSaturation: 0 <= value <= 1
type OxygenSaturationEmergencyFormSession struct {
	oxygenSaturation *float64
}

func NewOxygenSaturationEmergencyFormSession(oxygenSaturation float64) OxygenSaturationEmergencyFormSession {
	return OxygenSaturationEmergencyFormSession{oxygenSaturation: &oxygenSaturation}
}

func (f OxygenSaturationEmergencyFormSession) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.oxygenSaturation, ozzo.Min(0), ozzo.Max(1)),
	)
}

func (f OxygenSaturationEmergencyFormSession) IsSet() bool { return f.oxygenSaturation != nil }

func (f OxygenSaturationEmergencyFormSession) Score() float64 {
	saturarion := Undefined_OxygenSaturationLevel

	if !f.IsSet() {
		return saturarion.Float64()
	}

	switch {
	case *f.oxygenSaturation < 0.85:
		saturarion = SeverelyHypoxic_OxygenSaturationLevel
	case 0.85 <= *f.oxygenSaturation && *f.oxygenSaturation < 0.95:
		saturarion = Hypoxic_OxygenSaturationLevel
	case *f.oxygenSaturation >= 0.95:
		saturarion = Normal_OxygenSaturationLevel
	}

	return saturarion.Float64() / SeverelyHypoxic_OxygenSaturationLevel.Float64()
}
