package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbdominalPainIntensity_Float64(t *testing.T) {
	tests := []struct {
		name string
		i    AbdominalPainIntensity
		want float64
	}{
		{
			name: "Given a AbdominalPainIntensity, AbdominalPainIntensity.Float64() should return its value cast to float64",
			i:    VeryHigh_AbdominalPainIntensity,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.i.Float64())
		})
	}
}

func TestAbdominalPainIntensity_String(t *testing.T) {
	tests := []struct {
		name string
		i    AbdominalPainIntensity
		want string
	}{
		{
			name: "If AbdominalPainIntensity is VeryLow_AbdominalPainIntensity, AbdominalPainIntensity.String() should return 'very-low''",
			i:    VeryLow_AbdominalPainIntensity,
			want: "very-low",
		},
		{
			name: "If AbdominalPainIntensity is Low_AbdominalPainIntensity, AbdominalPainIntensity.String() should return 'low''",
			i:    Low_AbdominalPainIntensity,
			want: "low",
		},
		{
			name: "If AbdominalPainIntensity is Medium_AbdominalPainIntensity, AbdominalPainIntensity.String() should return 'medium''",
			i:    Medium_AbdominalPainIntensity,
			want: "medium",
		},
		{
			name: "If AbdominalPainIntensity is High_AbdominalPainIntensity, AbdominalPainIntensity.String() should return 'high''",
			i:    High_AbdominalPainIntensity,
			want: "high",
		},
		{
			name: "If AbdominalPainIntensity is VeryHigh_AbdominalPainIntensity, AbdominalPainIntensity.String() should return 'very-high''",
			i:    VeryHigh_AbdominalPainIntensity,
			want: "very-high",
		},
		{
			name: "If AbdominalPainIntensity isUndefined_AbdominalPainIntensity, AbdominalPainIntensity.String() should return 'undefined''",
			i:    Undefined_AbdominalPainIntensity,
			want: "undefined",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.i.String())
			})
		})
	}
}

func TestNewAbdominalPainIntensityFromString(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want AbdominalPainIntensity
	}{
		{
			name: "Given a string with value 'very-low', VeryLow_AbdominalPainIntensity should be returned",
			args: args{
				s: "very-low",
			},
			want: VeryLow_AbdominalPainIntensity,
		},
		{
			name: "Given a string with value 'low', Low_AbdominalPainIntensity should be returned",
			args: args{
				s: "low",
			},
			want: Low_AbdominalPainIntensity,
		},
		{
			name: "Given a string with value 'medium', Medium_AbdominalPainIntensity should be returned",
			args: args{
				s: "medium",
			},
			want: Medium_AbdominalPainIntensity,
		},
		{
			name: "Given a string with value 'high', High_AbdominalPainIntensity should be returned",
			args: args{
				s: "high",
			},
			want: High_AbdominalPainIntensity,
		},
		{
			name: "Given a string with value 'very-high', VeryHigh_AbdominalPainIntensity should be returned",
			args: args{
				s: "very-high",
			},
			want: VeryHigh_AbdominalPainIntensity,
		},
		{
			name: "Given a string unrelated to any AbdominalPainIntensity levels, Undefined_AbdominalPainIntensity should be returned",
			args: args{
				s: "banana",
			},
			want: Undefined_AbdominalPainIntensity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, NewAbdominalPainIntensityFromString(tt.args.s))
			})
		})
	}
}
