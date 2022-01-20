package valueobject

type EmergencyPriority int

const (
	Undefined_EmergencyPriority EmergencyPriority = iota
	Low_EmergencyPriority
	Medium_EmergencyPriority
	High_EmergencyPriority
	VeryHigh_EmergencyPriority
)

func (e EmergencyPriority) String() string {
	switch e {
	case Low_EmergencyPriority:
		return "low"
	case Medium_EmergencyPriority:
		return "medium"
	case High_EmergencyPriority:
		return "high"
	case VeryHigh_EmergencyPriority:
		return "very high"
	}

	return "undefined"
}
