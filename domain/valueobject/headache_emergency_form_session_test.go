package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeadacheEmergencyFormSession_Validate(t *testing.T) {
	type fields struct {
		Has       *bool
		Intensity HeadacheIntensity
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
			name: "If Has is nil, Intensity must be Undefined_HeadacheIntensity",
			fields: fields{
				Has:       nil,
				Intensity: High_HeadacheIntensity,
			},
			wantErr: true,
		},
		{
			name: "If Has is set as false, Intensity must be Undefined_HeadacheIntensity",
			fields: fields{
				Has:       &_false,
				Intensity: Low_HeadacheIntensity,
			},
			wantErr: true,
		},
		{
			name: "If Has is set as true, Intensity cannot be Undefined_HeadacheIntensity",
			fields: fields{
				Has:       &_true,
				Intensity: Undefined_HeadacheIntensity,
			},
			wantErr: true,
		},
		{
			name: "If Has is set as true and Intensity is not Undefined_HeadacheIntensity, Validate should return no errors",
			fields: fields{
				Has:       &_true,
				Intensity: Medium_HeadacheIntensity,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := HeadacheEmergencyFormSession{
				Has:       tt.fields.Has,
				Intensity: tt.fields.Intensity,
			}

			err := f.Validate()

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestHeadacheEmergencyFormSession_IsSet(t *testing.T) {
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
			f := HeadacheEmergencyFormSession{
				Has: tt.fields.Has,
			}

			assert.Equal(t, tt.want, f.IsSet())
		})
	}
}

func TestHeadacheEmergencyFormSession_Score(t *testing.T) {
	type fields struct {
		Intensity HeadacheIntensity
	}

	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Given an Intensity, HeadacheEmergencyFormSession.Score() should return its relative score towards a VeryHigh_HeadacheIntensity",
			fields: fields{
				Intensity: Medium_HeadacheIntensity,
			},
			want: 0.6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := HeadacheEmergencyFormSession{
				Intensity: tt.fields.Intensity,
			}

			assert.Equal(t, tt.want, f.Score())
		})
	}
}
