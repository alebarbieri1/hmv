package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

// UserData defines data about the User
type UserData struct {
	Name string
}

// Validate validates a UserData
func (d *UserData) Validate() error {
	return ozzo.ValidateStruct(d,
		ozzo.Field(&d.Name, ozzo.Required, ozzo.Length(0, 256)),
	)
}
