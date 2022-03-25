package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBodyTemperatureEmergencyFormSession_IsSet(t *testing.T) {
	type fields struct {
		CelsiusDegrees *float64
	}

	var (
		_celsiusDegrees = 36.0
	)

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "If CelsiusDegrees is not set, IsSet should return false",
			fields: fields{
				CelsiusDegrees: nil,
			},
			want: false,
		},
		{
			name: "If CelsiusDegrees is set, IsSet should return true",
			fields: fields{
				CelsiusDegrees: &_celsiusDegrees,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := BodyTemperatureEmergencyFormSession{
				CelsiusDegrees: tt.fields.CelsiusDegrees,
			}

			assert.Equal(t, tt.want, f.IsSet())
		})
	}
}

func TestBodyTemperatureEmergencyFormSession_Score(t *testing.T) {
	type fields struct {
		CelsiusDegrees *float64
	}

	var (
		t1 = 30.0
		t2 = 33.0
		t3 = 35.0
		t4 = 36.0
		t5 = 37.5
		t6 = 38.0
		t7 = 39.5
		t8 = 41.0
		t9 = 50.0
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
			name: "If CelsiusDegrees is 30.0, 1.0 should be returned",
			fields: fields{
				CelsiusDegrees: &t1,
			},
			want: 1.0,
		},
		{
			name: "If CelsiusDegrees is between 30.0 and 33.0, 0.75 should be returned",
			fields: fields{
				CelsiusDegrees: &t2,
			},
			want: 0.75,
		},
		{
			name: "If CelsiusDegrees is between 33.0 and 35.0, 0.5 should be returned",
			fields: fields{
				CelsiusDegrees: &t3,
			},
			want: 0.5,
		},
		{
			name: "If CelsiusDegrees is between 35.0 and 36.0, 0.25 should be returned",
			fields: fields{
				CelsiusDegrees: &t4,
			},
			want: 0.25,
		},
		{
			name: "If CelsiusDegrees is between 36.0 and 37.5, 0.0 should be returned",
			fields: fields{
				CelsiusDegrees: &t5,
			},
			want: 0.0,
		},
		{
			name: "If CelsiusDegrees is between 37.5 and 38.0, 0.25 should be returned",
			fields: fields{
				CelsiusDegrees: &t6,
			},
			want: 0.25,
		},
		{
			name: "If CelsiusDegrees is between 38.0 and 39.5, 0.5 should be returned",
			fields: fields{
				CelsiusDegrees: &t7,
			},
			want: 0.5,
		},
		{
			name: "If CelsiusDegrees is between 39.5 and 41.0, 0.75 should be returned",
			fields: fields{
				CelsiusDegrees: &t8,
			},
			want: 0.75,
		},
		{
			name: "If CelsiusDegrees is higher than 41.0, 1.0 should be returned",
			fields: fields{
				CelsiusDegrees: &t9,
			},
			want: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := BodyTemperatureEmergencyFormSession{
				CelsiusDegrees: tt.fields.CelsiusDegrees,
			}

			assert.Equal(t, tt.want, f.Score())
		})
	}
}
