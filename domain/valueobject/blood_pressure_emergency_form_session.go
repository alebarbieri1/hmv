package valueobject

var _ EmergencyFormSession = (*BloodPressureEmergencyFormSession)(nil)

type BloodPressureEmergencyFormSession struct {
	SystolicBloodPressure  *float64
	DiastolicBloodPressure *float64
}

func (f BloodPressureEmergencyFormSession) IsSet() bool {
	return f.SystolicBloodPressure != nil && f.DiastolicBloodPressure != nil
}

func (f BloodPressureEmergencyFormSession) Score() float64 {
	if f.IsSet() {
		switch {
		case *f.SystolicBloodPressure <= 105:
			return 0.8
		case 105 < *f.SystolicBloodPressure && *f.SystolicBloodPressure <= 120:
			return 0.0
		case 120 < *f.SystolicBloodPressure && *f.SystolicBloodPressure <= 130:
			return 0.2
		case 130 < *f.SystolicBloodPressure && *f.SystolicBloodPressure <= 140:
			return 0.4
		case 140 < *f.SystolicBloodPressure && *f.SystolicBloodPressure <= 160:
			return 0.6
		case 160 < *f.SystolicBloodPressure && *f.SystolicBloodPressure <= 180:
			return 0.8
		case *f.SystolicBloodPressure > 180:
			return 1.0
		}
	}

	return 0
}
