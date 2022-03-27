package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

// LocationData is data about a location
type LocationData struct {
	State   string
	City    string
	Address string
	ZipCode string
}

// Validate validates a LocationData
func (d *LocationData) Validate() error {
	return ozzo.ValidateStruct(d,
		ozzo.Field(&d.State, ozzo.Required, ozzo.Length(0, 256)),
		ozzo.Field(&d.City, ozzo.Required, ozzo.Length(0, 256)),
		ozzo.Field(&d.Address, ozzo.Required, ozzo.Length(0, 256)),
		ozzo.Field(&d.ZipCode, ozzo.Required, ozzo.Length(0, 256)),
	)
}
