package valueobject

type EmergencyForm struct {
	hasChestPain *bool
	hasHeadache  *bool
	priority     EmergencyPriority
}

func NewEmergencyForm(hasChestPain, hasHeadache *bool) EmergencyForm {
	f := EmergencyForm{
		hasChestPain: hasChestPain,
		hasHeadache:  hasHeadache,
	}

	f.priority = f.recalculatePriority()

	return f
}

func (f EmergencyForm) HasChestPain() *bool { return f.hasChestPain }

func (f EmergencyForm) HasHeadache() *bool { return f.hasHeadache }

func (f EmergencyForm) Priority() EmergencyPriority { return f.priority }

func (f EmergencyForm) isComplete() bool {
	if f.hasChestPain == nil {
		return false
	}

	if f.hasHeadache == nil {
		return false
	}

	return true
}

func (f EmergencyForm) recalculatePriority() EmergencyPriority {
	if !f.isComplete() {
		return Undefined_EmergencyPriority
	}

	switch {
	case *f.hasChestPain && *f.hasHeadache:
		return VeryHigh_EmergencyPriority
	case *f.hasChestPain:
		return High_EmergencyPriority
	case *f.hasHeadache:
		return Medium_EmergencyPriority
	}

	return Low_EmergencyPriority
}
