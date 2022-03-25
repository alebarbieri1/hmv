package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbdominalPainEmergencyFormSession_Validate(t *testing.T) {
	type fields struct {
		Has       *bool
		Intensity AbdominalPainIntensity
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
			name: "If Has is nil, Intensity must be Undefined_AbdominalPainIntensity",
			fields: fields{
				Has:       nil,
				Intensity: High_AbdominalPainIntensity,
			},
			wantErr: true,
		},
		{
			name: "If Has is set as false, Intensity must be Undefined_AbdominalPainIntensity",
			fields: fields{
				Has:       &_false,
				Intensity: Low_AbdominalPainIntensity,
			},
			wantErr: true,
		},
		{
			name: "If Has is set as true, Intensity cannot be Undefined_AbdominalPainIntensity",
			fields: fields{
				Has:       &_true,
				Intensity: Undefined_AbdominalPainIntensity,
			},
			wantErr: true,
		},
		{
			name: "If Has is set as true and Intensity is not Undefined_AbdominalPainIntensity, Validate should return no errors",
			fields: fields{
				Has:       &_true,
				Intensity: Medium_AbdominalPainIntensity,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := AbdominalPainEmergencyFormSession{
				Has:       tt.fields.Has,
				Intensity: tt.fields.Intensity,
			}

			err := f.Validate()

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestAbdominalPainEmergencyFormSession_IsSet(t *testing.T) {
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
			f := AbdominalPainEmergencyFormSession{
				Has: tt.fields.Has,
			}

			assert.Equal(t, tt.want, f.IsSet())
		})
	}
}

func TestAbdominalPainEmergencyFormSession_Score(t *testing.T) {
	type fields struct {
		Intensity AbdominalPainIntensity
	}

	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Given an Intensity, AbdominalPainEmergencyFormSession.Score() should return its relative score towards a VeryHigh_AbdominalPainIntensity",
			fields: fields{
				Intensity: Medium_AbdominalPainIntensity,
			},
			want: 0.6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := AbdominalPainEmergencyFormSession{
				Intensity: tt.fields.Intensity,
			}

			assert.Equal(t, tt.want, f.Score())
		})
	}
}
