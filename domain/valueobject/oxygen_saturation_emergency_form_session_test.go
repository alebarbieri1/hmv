package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOxygenSaturationEmergencyFormSession_Validate(t *testing.T) {
	type fields struct {
		Value *float64
	}

	var (
		v1 = -1.0
		v2 = 2.0
		v3 = 0.5
	)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "If Value is nil, Validate should return no errors",
			fields: fields{
				Value: nil,
			},
			wantErr: false,
		},
		{
			name: "If Value is set with a negative value, an error should be returned",
			fields: fields{
				Value: &v1,
			},
			wantErr: true,
		},
		{
			name: "If Value is set with a value higher than 1.0, an error should be returned",
			fields: fields{
				Value: &v2,
			},
			wantErr: true,
		},
		{
			name: "If Value is set with a value between 0.0 and 1.0, Validate should return no errors",
			fields: fields{
				Value: &v3,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := OxygenSaturationEmergencyFormSession{
				Value: tt.fields.Value,
			}

			err := f.Validate()

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestOxygenSaturationEmergencyFormSession_IsSet(t *testing.T) {
	type fields struct {
		Value *float64
	}

	value := 0.5

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "If Value is not set, IsSet should return false",
			fields: fields{
				Value: nil,
			},
			want: false,
		},
		{
			name: "If Value is set with any values, IsSet should return true",
			fields: fields{
				Value: &value,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := OxygenSaturationEmergencyFormSession{
				Value: tt.fields.Value,
			}

			assert.Equal(t, tt.want, f.IsSet())
		})
	}
}

func TestOxygenSaturationEmergencyFormSession_Score(t *testing.T) {
	type fields struct {
		Value *float64
	}

	var (
		v1 = 0.84
		v2 = 0.94
		v3 = 0.95
	)

	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "If the session is not set, 0.0 should be returned",
			fields: fields{},
			want:   0.0,
		},
		{
			name: "If Value is 0.85 or lower, 1.0 should be returned",
			fields: fields{
				Value: &v1,
			},
			want: 1.0,
		},
		{
			name: "If Value is between 0.85 and 0.95, 0.75 should be returned",
			fields: fields{
				Value: &v2,
			},
			want: 0.75,
		},
		{
			name: "If Value is higher than 0.95, 0.0 should be returned",
			fields: fields{
				Value: &v3,
			},
			want: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := OxygenSaturationEmergencyFormSession{
				Value: tt.fields.Value,
			}

			assert.Equal(t, tt.want, f.Score())
		})
	}
}
