package presenter

import "flavioltonon/hmv/domain/valueobject"

type EmergencyContact struct {
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
}

func NewEmergencyContact(c valueobject.EmergencyContact) *EmergencyContact {
	return &EmergencyContact{
		Name:         c.Name,
		MobileNumber: c.MobileNumber,
	}
}
