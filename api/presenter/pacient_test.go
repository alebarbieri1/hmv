package presenter

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPacient(t *testing.T) {
	type args struct {
		e *entity.Pacient
	}

	date := time.Date(2022, time.January, 25, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		args args
		want *Pacient
	}{
		{
			name: "Given a Pacient, a valid presentation should be returned",
			args: args{
				e: &entity.Pacient{
					ID:     "1",
					UserID: "2",
					Data: valueobject.PacientData{
						Name:      "foo",
						BirthDate: date,
						Location: valueobject.LocationData{
							State:   "bar",
							City:    "baz",
							Address: "qux",
							ZipCode: "3",
						},
						EmergencyContact: valueobject.EmergencyContact{
							Name:         "foo",
							MobileNumber: "bar",
						},
						Health: valueobject.HealthData{
							AllergicToMedicationDescription:          "foo1",
							ContinuousUseMedicationDescription:       "foo2",
							DiagnosedDiseaseDescription:              "foo3",
							HistoryOfHeartDiseaseDescription:         "foo4",
							SmokingBehaviorDescription:               "foo5",
							PhysicalActivityDescription:              "foo6",
							AlcoholicBeveragesConsumptionDescription: "foo7",
						},
					},
					CreatedAt: date,
					UpdatedAt: date,
				},
			},
			want: &Pacient{
				ID:        "1",
				UserID:    "2",
				Name:      "foo",
				BirthDate: "25/01/2022 - 00:00:00h",
				Location: &LocationData{
					State:   "bar",
					City:    "baz",
					Address: "qux",
					ZipCode: "3",
				},
				EmergencyContact: &EmergencyContact{
					Name:         "foo",
					MobileNumber: "bar",
				},
				Health: &HealthData{
					AllergicToMedicationDescription:          "foo1",
					ContinuousUseMedicationDescription:       "foo2",
					DiagnosedDiseaseDescription:              "foo3",
					HistoryOfHeartDiseaseDescription:         "foo4",
					SmokingBehaviorDescription:               "foo5",
					PhysicalActivityDescription:              "foo6",
					AlcoholicBeveragesConsumptionDescription: "foo7",
				},
				CreatedAt: "25/01/2022 - 00:00:00h",
				UpdatedAt: "25/01/2022 - 00:00:00h",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewPacient(tt.args.e))
		})
	}
}
