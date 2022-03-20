package valueobject

var _ EmergencyFormSession = (*BodyTemperatureEmergencyFormSession)(nil)

type BodyTemperatureEmergencyFormSession struct {
	CelsiusDegrees *float64
}

func (f BodyTemperatureEmergencyFormSession) IsSet() bool { return f.CelsiusDegrees != nil }

func (f BodyTemperatureEmergencyFormSession) Score() float64 {
	if f.IsSet() {
		switch {
		case *f.CelsiusDegrees <= 30.0:
			return 1.0
		case 30.0 < *f.CelsiusDegrees && *f.CelsiusDegrees <= 33.0:
			return 0.75
		case 33.0 < *f.CelsiusDegrees && *f.CelsiusDegrees <= 35.0:
			return 0.50
		case 35.0 < *f.CelsiusDegrees && *f.CelsiusDegrees <= 36.0:
			return 0.25
		case 37.5 < *f.CelsiusDegrees && *f.CelsiusDegrees <= 38.0:
			return 0.25
		case 38.0 < *f.CelsiusDegrees && *f.CelsiusDegrees <= 39.5:
			return 0.50
		case 39.5 < *f.CelsiusDegrees && *f.CelsiusDegrees <= 41.0:
			return 0.75
		case *f.CelsiusDegrees > 41.0:
			return 1.0
		}
	}

	return 0
}
