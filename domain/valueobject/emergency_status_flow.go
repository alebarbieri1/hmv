package valueobject

// EmergencyStatusFlow defines flows that are allowed on EmergencyStatus changes
type EmergencyStatusFlow map[EmergencyStatus][]EmergencyStatus

// CanChange returns true if a given EmergencyStatus (from) can progress to another given EmergencyStatus (to)
func (f EmergencyStatusFlow) CanChange(from, to EmergencyStatus) bool {
	fromStatusList, exists := f[from]
	if !exists {
		return false
	}

	for _, status := range fromStatusList {
		if status == to {
			return true
		}
	}

	return false
}

// DefaultEmergencyStatusFlow defines all the default EmergencyStatus changes available
var DefaultEmergencyStatusFlow = EmergencyStatusFlow{
	Undefined_EmergencyStatus: {
		Triage_EmergencyStatus,
		AmbulanceToPacient_EmergencyStatus,
		AmbulanceToHospital_EmergencyStatus,
		Finished_EmergencyStatus,
		Cancelled_EmergencyStatus,
	},
	Triage_EmergencyStatus: {
		AmbulanceToPacient_EmergencyStatus,
		Cancelled_EmergencyStatus,
	},
	AmbulanceToPacient_EmergencyStatus: {
		AmbulanceToHospital_EmergencyStatus,
		Cancelled_EmergencyStatus,
	},
	AmbulanceToHospital_EmergencyStatus: {
		Finished_EmergencyStatus,
		Cancelled_EmergencyStatus,
	},
}
