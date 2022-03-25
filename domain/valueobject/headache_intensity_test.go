package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeadacheIntensity_Float64(t *testing.T) {
	tests := []struct {
		name string
		i    HeadacheIntensity
		want float64
	}{
		{
			name: "Given a HeadacheIntensity, HeadacheIntensity.Float64() should return its value cast to float64",
			i:    VeryHigh_HeadacheIntensity,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.i.Float64())
		})
	}
}

func TestHeadacheIntensity_String(t *testing.T) {
	tests := []struct {
		name string
		i    HeadacheIntensity
		want string
	}{
		{
			name: "If HeadacheIntensity is VeryLow_HeadacheIntensity, HeadacheIntensity.String() should return 'very-low''",
			i:    VeryLow_HeadacheIntensity,
			want: "very-low",
		},
		{
			name: "If HeadacheIntensity is Low_HeadacheIntensity, HeadacheIntensity.String() should return 'low''",
			i:    Low_HeadacheIntensity,
			want: "low",
		},
		{
			name: "If HeadacheIntensity is Medium_HeadacheIntensity, HeadacheIntensity.String() should return 'medium''",
			i:    Medium_HeadacheIntensity,
			want: "medium",
		},
		{
			name: "If HeadacheIntensity is High_HeadacheIntensity, HeadacheIntensity.String() should return 'high''",
			i:    High_HeadacheIntensity,
			want: "high",
		},
		{
			name: "If HeadacheIntensity is VeryHigh_HeadacheIntensity, HeadacheIntensity.String() should return 'very-high''",
			i:    VeryHigh_HeadacheIntensity,
			want: "very-high",
		},
		{
			name: "If HeadacheIntensity is Undefined_HeadacheIntensity, HeadacheIntensity.String() should return 'undefined''",
			i:    Undefined_HeadacheIntensity,
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

func TestNewHeadacheIntensityFromString(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want HeadacheIntensity
	}{
		{
			name: "Given a string with value 'very-low', VeryLow_HeadacheIntensity should be returned",
			args: args{
				s: "very-low",
			},
			want: VeryLow_HeadacheIntensity,
		},
		{
			name: "Given a string with value 'low', Low_HeadacheIntensity should be returned",
			args: args{
				s: "low",
			},
			want: Low_HeadacheIntensity,
		},
		{
			name: "Given a string with value 'medium', Medium_HeadacheIntensity should be returned",
			args: args{
				s: "medium",
			},
			want: Medium_HeadacheIntensity,
		},
		{
			name: "Given a string with value 'high', High_HeadacheIntensity should be returned",
			args: args{
				s: "high",
			},
			want: High_HeadacheIntensity,
		},
		{
			name: "Given a string with value 'very-high', VeryHigh_HeadacheIntensity should be returned",
			args: args{
				s: "very-high",
			},
			want: VeryHigh_HeadacheIntensity,
		},
		{
			name: "Given a string unrelated to any HeadacheIntensity levels, Undefined_HeadacheIntensity should be returned",
			args: args{
				s: "banana",
			},
			want: Undefined_HeadacheIntensity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, NewHeadacheIntensityFromString(tt.args.s))
			})
		})
	}
}
