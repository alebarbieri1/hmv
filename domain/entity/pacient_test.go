package entity

import (
	"flavioltonon/hmv/domain/valueobject"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPacient_Validate(t *testing.T) {
	type fields struct {
		ID        string
		UserID    string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	var (
		today    = time.Now().Truncate(24 * time.Hour)
		tomorrow = today.Add(24 * time.Hour)
	)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "If all fields in Pacient are valid, Pacient.Validate() should return no errors",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
			wantErr: false,
		},
		{
			name: "If ID is empty, Pacient.Validate() should return an error",
			fields: fields{
				ID:        "",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
			wantErr: true,
		},
		{
			name: "If ID contains a non-UUIDv4 value, Pacient.Validate() should return an error",
			fields: fields{
				ID:        "foo",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
			wantErr: true,
		},
		{
			name: "If UserID is empty, Pacient.Validate() should return an error",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "",
				CreatedAt: today,
				UpdatedAt: today,
			},
			wantErr: true,
		},
		{
			name: "If UserID contains a non-UUIDv4 value, Pacient.Validate() should return an error",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "foo",
				CreatedAt: today,
				UpdatedAt: today,
			},
			wantErr: true,
		},
		{
			name: "If CreatedAt contains data in the future, Pacient.Validate() should return an error",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: tomorrow,
				UpdatedAt: today,
			},
			wantErr: true,
		},
		{
			name: "If UpdatedAt contains data in the future, Pacient.Validate() should return an error",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: tomorrow,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pacient{
				ID:        tt.fields.ID,
				UserID:    tt.fields.UserID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}

			err := p.Validate()

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestPacient_UpdateEmergencyContact(t *testing.T) {
	type fields struct {
		EmergencyContact valueobject.EmergencyContact
	}

	type args struct {
		emergencyContact valueobject.EmergencyContact
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    *Pacient
	}{
		{
			name: "Given a valid valueobject.EmergencyContact, Pacient.EmergencyContact should be updated",
			fields: fields{
				EmergencyContact: valueobject.EmergencyContact{
					Name:         "foo",
					MobileNumber: "5511999999999",
				},
			},
			args: args{
				emergencyContact: valueobject.EmergencyContact{
					Name:         "bar",
					MobileNumber: "5519999999999",
				},
			},
			wantErr: false,
			want: &Pacient{
				Data: valueobject.PacientData{
					EmergencyContact: valueobject.EmergencyContact{
						Name:         "bar",
						MobileNumber: "5519999999999",
					},
				},
			},
		},
		{
			name: "Given an invalid valueobject.EmergencyContact, Pacient.EmergencyContact should not be updated",
			fields: fields{
				EmergencyContact: valueobject.EmergencyContact{
					Name:         "foo",
					MobileNumber: "5511999999999",
				},
			},
			args: args{
				emergencyContact: valueobject.EmergencyContact{
					Name:         "",
					MobileNumber: "5519999999999",
				},
			},
			wantErr: true,
			want: &Pacient{
				Data: valueobject.PacientData{
					EmergencyContact: valueobject.EmergencyContact{
						Name:         "foo",
						MobileNumber: "5511999999999",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pacient{
				Data: valueobject.PacientData{
					EmergencyContact: tt.fields.EmergencyContact,
				},
			}

			err := p.UpdateEmergencyContact(tt.args.emergencyContact)

			assert.Equal(t, tt.wantErr, err != nil)

			if err == nil {
				assert.Equal(t, tt.want.Data, p.Data)
			}
		})
	}
}
