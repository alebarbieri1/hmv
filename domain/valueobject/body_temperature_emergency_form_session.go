package valueobject

var _ EmergencyFormSession = (*BodyTemperatureEmergencyFormSession)(nil)

type BodyTemperatureEmergencyFormSession struct {
	bodyTemperature *float64
}

func NewBodyTemperatureEmergencyFormSession(bodyTemperature float64) BodyTemperatureEmergencyFormSession {
	return BodyTemperatureEmergencyFormSession{bodyTemperature: &bodyTemperature}
}

func (f BodyTemperatureEmergencyFormSession) IsSet() bool { return f.bodyTemperature != nil }

func (f BodyTemperatureEmergencyFormSession) Score() float64 {
	if f.IsSet() {
		switch {
		case *f.bodyTemperature <= 30.0:
			return 1.0
		case 30.0 < *f.bodyTemperature && *f.bodyTemperature <= 33.0:
			return 0.75
		case 33.0 < *f.bodyTemperature && *f.bodyTemperature <= 35.0:
			return 0.50
		case 35.0 < *f.bodyTemperature && *f.bodyTemperature <= 36.0:
			return 0.25
		case 37.5 < *f.bodyTemperature && *f.bodyTemperature <= 38.0:
			return 0.25
		case 38.0 < *f.bodyTemperature && *f.bodyTemperature <= 39.5:
			return 0.50
		case 39.5 < *f.bodyTemperature && *f.bodyTemperature <= 41.0:
			return 0.75
		case *f.bodyTemperature > 41.0:
			return 1.0
		}
	}

	return 0
}
