package valueobject

var _ EmergencyFormSession = (*BloodPressureEmergencyFormSession)(nil)

type BloodPressureEmergencyFormSession struct {
	Systolic  *float64
	Diastolic *float64
}

func (f BloodPressureEmergencyFormSession) IsSet() bool {
	return f.Systolic != nil && f.Diastolic != nil
}

func (f BloodPressureEmergencyFormSession) Score() float64 {
	if f.IsSet() {
		switch {
		case *f.Systolic <= 105:
			return 0.8
		case 105 < *f.Systolic && *f.Systolic <= 120:
			return 0.0
		case 120 < *f.Systolic && *f.Systolic <= 130:
			return 0.2
		case 130 < *f.Systolic && *f.Systolic <= 140:
			return 0.4
		case 140 < *f.Systolic && *f.Systolic <= 160:
			return 0.6
		case 160 < *f.Systolic && *f.Systolic <= 180:
			return 0.8
		case *f.Systolic > 180:
			return 1.0
		}
	}

	return 0
}
