package valueobject

type EmergencyStatus int

const (
	Undefined_EmergencyStatus EmergencyStatus = iota
	Triage_EmergencyStatus
	AmbulanceToPacient_EmergencyStatus
	AmbulanceToHospital_EmergencyStatus
	Finished_EmergencyStatus
	Cancelled_EmergencyStatus
)

func NewEmergencyStatusFromString(s string) EmergencyStatus {
	switch s {
	case "triage":
		return Triage_EmergencyStatus
	case "ambulance-to-pacient":
		return AmbulanceToPacient_EmergencyStatus
	case "ambulance-to-hospital":
		return AmbulanceToHospital_EmergencyStatus
	case "finished":
		return Finished_EmergencyStatus
	case "cancelled":
		return Cancelled_EmergencyStatus
	}

	return Undefined_EmergencyStatus
}

func (s EmergencyStatus) String() string {
	switch s {
	case Triage_EmergencyStatus:
		return "triage"
	case AmbulanceToPacient_EmergencyStatus:
		return "ambulance-to-pacient"
	case AmbulanceToHospital_EmergencyStatus:
		return "ambulance-to-hospital"
	case Finished_EmergencyStatus:
		return "finished"
	case Cancelled_EmergencyStatus:
		return "cancelled"
	}

	return "undefined"
}
