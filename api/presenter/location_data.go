package presenter

import "flavioltonon/hmv/domain/valueobject"

type LocationData struct {
	State   string `json:"state"`
	City    string `json:"city"`
	Address string `json:"address"`
	ZipCode string `json:"zipcode"`
}

func NewLocationData(data valueobject.LocationData) *LocationData {
	return &LocationData{
		State:   data.State,
		City:    data.City,
		Address: data.Address,
		ZipCode: data.ZipCode,
	}
}
