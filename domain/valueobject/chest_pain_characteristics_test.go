package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChestPainCharacteristics_Float64(t *testing.T) {
	tests := []struct {
		name string
		i    ChestPainCharacteristics
		want float64
	}{
		{
			name: "Given a ChestPainCharacteristics, ChestPainCharacteristics.Float64() should return its value cast to float64",
			i:    RadiatingToTheLeftArm_ChestPainCharacteristics,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.i.Float64())
		})
	}
}

func TestChestPainCharacteristics_String(t *testing.T) {
	tests := []struct {
		name string
		i    ChestPainCharacteristics
		want string
	}{
		{
			name: "If ChestPainCharacteristics is RadiatingToTheLeftArm_ChestPainCharacteristics, ChestPainCharacteristics.String() should return 'radiating-to-the-left-arm''",
			i:    RadiatingToTheLeftArm_ChestPainCharacteristics,
			want: "radiating-to-the-left-arm",
		},
		{
			name: "If ChestPainCharacteristics is Undefined_ChestPainCharacteristics, ChestPainCharacteristics.String() should return 'undefined''",
			i:    Undefined_ChestPainCharacteristics,
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

func TestNewChestPainCharacteristicsFromString(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want ChestPainCharacteristics
	}{
		{
			name: "Given a string with value 'very-low', RadiatingToTheLeftArm_ChestPainCharacteristics should be returned",
			args: args{
				s: "radiating-to-the-left-arm",
			},
			want: RadiatingToTheLeftArm_ChestPainCharacteristics,
		},
		{
			name: "Given a string unrelated to any ChestPainCharacteristics levels, Undefined_ChestPainCharacteristics should be returned",
			args: args{
				s: "banana",
			},
			want: Undefined_ChestPainCharacteristics,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, NewChestPainCharacteristicsFromString(tt.args.s))
			})
		})
	}
}
