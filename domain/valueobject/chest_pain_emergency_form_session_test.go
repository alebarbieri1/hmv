package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChestPainEmergencyFormSession_Validate(t *testing.T) {
	type fields struct {
		Has             *bool
		Characteristics ChestPainCharacteristics
	}

	var (
		_true  = true
		_false = false
	)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "If Has is nil, Characteristics must be Undefined_ChestPainCharacteristics",
			fields: fields{
				Has:             nil,
				Characteristics: RadiatingToTheLeftArm_ChestPainCharacteristics,
			},
			wantErr: true,
		},
		{
			name: "If Has is set as false, Characteristics must be Undefined_ChestPainCharacteristics",
			fields: fields{
				Has:             &_false,
				Characteristics: RadiatingToTheLeftArm_ChestPainCharacteristics,
			},
			wantErr: true,
		},
		{
			name: "If Has is set as true, Characteristics cannot be Undefined_ChestPainCharacteristics",
			fields: fields{
				Has:             &_true,
				Characteristics: Undefined_ChestPainCharacteristics,
			},
			wantErr: true,
		},
		{
			name: "If Has is set as true and Characteristics is not Undefined_ChestPainCharacteristics, Validate should return no errors",
			fields: fields{
				Has:             &_true,
				Characteristics: RadiatingToTheLeftArm_ChestPainCharacteristics,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := ChestPainEmergencyFormSession{
				Has:             tt.fields.Has,
				Characteristics: tt.fields.Characteristics,
			}

			err := f.Validate()

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestChestPainEmergencyFormSession_IsSet(t *testing.T) {
	type fields struct {
		Has *bool
	}

	_false := false

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "If Has is not set, IsSet should return false",
			fields: fields{
				Has: nil,
			},
			want: false,
		},
		{
			name: "If Has is set with any values, IsSet should return true",
			fields: fields{
				Has: &_false,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := ChestPainEmergencyFormSession{
				Has: tt.fields.Has,
			}

			assert.Equal(t, tt.want, f.IsSet())
		})
	}
}

func TestChestPainEmergencyFormSession_Score(t *testing.T) {
	type fields struct {
		Characteristics ChestPainCharacteristics
	}

	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Given an Characteristics, ChestPainEmergencyFormSession.Score() should return its relative score towards a RadiatingToTheLeftArm_ChestPainCharacteristics",
			fields: fields{
				Characteristics: RadiatingToTheLeftArm_ChestPainCharacteristics,
			},
			want: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := ChestPainEmergencyFormSession{
				Characteristics: tt.fields.Characteristics,
			}

			assert.Equal(t, tt.want, f.Score())
		})
	}
}
