package presenter

import "flavioltonon/hmv/domain/valueobject"

// EmergencyContact is a valueobject.EmergencyContact presenter
type EmergencyContact struct {
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
}

// NewEmergencyContact returns a presentation for a EmergencyContact
func NewEmergencyContact(c valueobject.EmergencyContact) *EmergencyContact {
	return &EmergencyContact{
		Name:         c.Name,
		MobileNumber: c.MobileNumber,
	}
}
