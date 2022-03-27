package valueobject

import (
	"time"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

// PacientData defines data about the Pacient
type PacientData struct {
	Name             string
	BirthDate        time.Time
	Location         LocationData
	EmergencyContact EmergencyContact
	Health           HealthData
}

// Validate validates a PacientData
func (d *PacientData) Validate() error {
	today := time.Now().Truncate(24 * time.Hour)

	return ozzo.ValidateStruct(d,
		ozzo.Field(&d.Name, ozzo.Required, ozzo.Length(0, 256)),
		ozzo.Field(&d.BirthDate, ozzo.Required, ozzo.Max(today)),
		ozzo.Field(&d.Location),
		ozzo.Field(&d.EmergencyContact),
		ozzo.Field(&d.Health),
	)
}
