package presenter

import "flavioltonon/hmv/domain/valueobject"

type EmergencyForm struct {
	HasChestPain *bool  `json:"has_chest_pain"`
	HasHeadache  *bool  `json:"has_headache"`
	Priority     string `json:"priority"`
}

func NewEmergencyForm(f valueobject.EmergencyForm) *EmergencyForm {
	return &EmergencyForm{
		HasChestPain: f.HasChestPain(),
		HasHeadache:  f.HasHeadache(),
		Priority:     f.Priority().String(),
	}
}
