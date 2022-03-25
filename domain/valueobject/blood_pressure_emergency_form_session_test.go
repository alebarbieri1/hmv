package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBloodPressureEmergencyFormSession_IsSet(t *testing.T) {
	type fields struct {
		Systolic  *float64
		Diastolic *float64
	}

	var (
		_systolic  = 120.0
		_diastolic = 80.0
	)

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "If Systolic and Diastolic are not set, IsSet should return false",
			fields: fields{
				Systolic:  nil,
				Diastolic: nil,
			},
			want: false,
		},
		{
			name: "If Systolic is set, but Diastolic is not, IsSet should return false",
			fields: fields{
				Systolic:  &_systolic,
				Diastolic: nil,
			},
			want: false,
		},
		{
			name: "If Systolic is not set and Diastolic is set, IsSet should return false",
			fields: fields{
				Systolic:  nil,
				Diastolic: &_diastolic,
			},
			want: false,
		},
		{
			name: "If both Systolic and Diastolic are set, IsSet should return true",
			fields: fields{
				Systolic:  &_systolic,
				Diastolic: &_diastolic,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := BloodPressureEmergencyFormSession{
				Systolic:  tt.fields.Systolic,
				Diastolic: tt.fields.Diastolic,
			}

			assert.Equal(t, tt.want, f.IsSet())
		})
	}
}

func TestBloodPressureEmergencyFormSession_Score(t *testing.T) {
	type fields struct {
		Systolic  *float64
		Diastolic *float64
	}

	var (
		p1 = 105.0
		p2 = 120.0
		p3 = 130.0
		p4 = 140.0
		p5 = 160.0
		p6 = 180.0
		p7 = 200.0
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
			name: "If Systolic is 105 or lower, 0.8 should be returned",
			fields: fields{
				Systolic:  &p1,
				Diastolic: &p1,
			},
			want: 0.8,
		},
		{
			name: "If Systolic is between 105 and 120, 0.0 should be returned",
			fields: fields{
				Systolic:  &p2,
				Diastolic: &p1,
			},
			want: 0.0,
		},
		{
			name: "If Systolic is between 120 and 130, 0.2 should be returned",
			fields: fields{
				Systolic:  &p3,
				Diastolic: &p1,
			},
			want: 0.2,
		},
		{
			name: "If Systolic is between 130 and 140, 0.4 should be returned",
			fields: fields{
				Systolic:  &p4,
				Diastolic: &p1,
			},
			want: 0.4,
		},
		{
			name: "If Systolic is between 140 and 160, 0.6 should be returned",
			fields: fields{
				Systolic:  &p5,
				Diastolic: &p1,
			},
			want: 0.6,
		},
		{
			name: "If Systolic is between 160 and 180, 0.8 should be returned",
			fields: fields{
				Systolic:  &p6,
				Diastolic: &p1,
			},
			want: 0.8,
		},
		{
			name: "If Systolic is higher than 180, 1.0 should be returned",
			fields: fields{
				Systolic:  &p7,
				Diastolic: &p1,
			},
			want: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := BloodPressureEmergencyFormSession{
				Systolic:  tt.fields.Systolic,
				Diastolic: tt.fields.Diastolic,
			}

			assert.Equal(t, tt.want, f.Score())
		})
	}
}
