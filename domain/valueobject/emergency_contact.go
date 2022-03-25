package valueobject

import (
	"github.com/flavioltonon/go-brazil"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

// EmergencyContact represents an emergency contact of a Pacient
type EmergencyContact struct {
	Name         string
	MobileNumber string
}

// Validate validates the EmergencyContact
func (i EmergencyContact) Validate() error {
	return ozzo.ValidateStruct(&i,
		ozzo.Field(&i.Name, ozzo.Required),
		ozzo.Field(&i.MobileNumber, ozzo.Required, ozzo.By(i.validateMobileNumber)),
	)
}

// validateMobileNumber validates a given value with a mobile number validator
func (i EmergencyContact) validateMobileNumber(value interface{}) error {
	v, err := ozzo.EnsureString(value)
	if err != nil {
		return err
	}

	if _, err := brazil.ParseMobile(v); err != nil {
		return err
	}

	return nil
}
