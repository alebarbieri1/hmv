package valueobject

var _ EmergencyFormSession = (*BodyTemperatureEmergencyFormSession)(nil)

type BodyTemperatureEmergencyFormSession struct {
	BodyTemperature *float64
}

func (f BodyTemperatureEmergencyFormSession) IsSet() bool { return f.BodyTemperature != nil }

func (f BodyTemperatureEmergencyFormSession) Score() float64 {
	if f.IsSet() {
		switch {
		case *f.BodyTemperature <= 30.0:
			return 1.0
		case 30.0 < *f.BodyTemperature && *f.BodyTemperature <= 33.0:
			return 0.75
		case 33.0 < *f.BodyTemperature && *f.BodyTemperature <= 35.0:
			return 0.50
		case 35.0 < *f.BodyTemperature && *f.BodyTemperature <= 36.0:
			return 0.25
		case 37.5 < *f.BodyTemperature && *f.BodyTemperature <= 38.0:
			return 0.25
		case 38.0 < *f.BodyTemperature && *f.BodyTemperature <= 39.5:
			return 0.50
		case 39.5 < *f.BodyTemperature && *f.BodyTemperature <= 41.0:
			return 0.75
		case *f.BodyTemperature > 41.0:
			return 1.0
		}
	}

	return 0
}
