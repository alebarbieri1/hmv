package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmergencyContact_validateMobileNumber(t *testing.T) {
	type args struct {
		value interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Given a non-string value, EmergencyContact.validateMobileNumber should return an error",
			args: args{
				value: 123,
			},
			wantErr: true,
		},
		{
			name: "Given an invalid mobile number string value, EmergencyContact.validateMobileNumber should return an error",
			args: args{
				value: "123456",
			},
			wantErr: true,
		},
		{
			name: "Given a valid mobile number string value, EmergencyContact.validateMobileNumber should return no errors",
			args: args{
				value: "5511999999999",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := EmergencyContact{}

			err := i.validateMobileNumber(tt.args.value)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestEmergencyContact_Validate(t *testing.T) {
	type fields struct {
		Name         string
		MobileNumber string
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "If all fields are valid, EmergencyContact.Validate() should return no errors",
			fields: fields{
				Name:         "foo",
				MobileNumber: "5511999999999",
			},
			wantErr: false,
		},
		{
			name: "If no Name is provided, EmergencyContact.Validate() should return an error",
			fields: fields{
				MobileNumber: "5511999999999",
			},
			wantErr: true,
		},
		{
			name: "If no MobileNumber is provided, EmergencyContact.Validate() should return an error",
			fields: fields{
				Name: "foo",
			},
			wantErr: true,
		},
		{
			name: "If an invalid MobileNumber is provided, EmergencyContact.Validate() should return an error",
			fields: fields{
				Name:         "foo",
				MobileNumber: "bar",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := EmergencyContact{
				Name:         tt.fields.Name,
				MobileNumber: tt.fields.MobileNumber,
			}

			err := i.Validate()

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
