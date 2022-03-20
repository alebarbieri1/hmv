package valueobject

var _ EmergencyFormSession = (*BloodPressureEmergencyFormSession)(nil)

type BloodPressureEmergencyFormSession struct {
	systolicBloodPressure  *float64
	diastolicBloodPressure *float64
}

func NewBloodPressureEmergencyFormSession(systolicBloodPressure, diastolicBloodPressure float64) BloodPressureEmergencyFormSession {
	return BloodPressureEmergencyFormSession{systolicBloodPressure: &systolicBloodPressure, diastolicBloodPressure: &diastolicBloodPressure}
}

func (f BloodPressureEmergencyFormSession) IsSet() bool {
	return f.systolicBloodPressure != nil && f.diastolicBloodPressure != nil
}

func (f BloodPressureEmergencyFormSession) Score() float64 {
	if f.IsSet() {
		switch {
		case *f.systolicBloodPressure <= 105:
			return 0.8
		case 105 < *f.systolicBloodPressure && *f.systolicBloodPressure <= 120:
			return 0.0
		case 120 < *f.systolicBloodPressure && *f.systolicBloodPressure <= 130:
			return 0.2
		case 130 < *f.systolicBloodPressure && *f.systolicBloodPressure <= 140:
			return 0.4
		case 140 < *f.systolicBloodPressure && *f.systolicBloodPressure <= 160:
			return 0.6
		case 160 < *f.systolicBloodPressure && *f.systolicBloodPressure <= 180:
			return 0.8
		case *f.systolicBloodPressure > 180:
			return 1.0
		}
	}

	return 0
}
