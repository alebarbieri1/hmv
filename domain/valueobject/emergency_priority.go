package valueobject

// EmergencyPriority is the priority level given to an Emergency
type EmergencyPriority int

const (
	Undefined_EmergencyPriority EmergencyPriority = iota
	Low_EmergencyPriority
	Medium_EmergencyPriority
	High_EmergencyPriority
	VeryHigh_EmergencyPriority
)

// String returns the string value of an EmergencyPriority
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
