package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOxygenSaturationLevel_Float64(t *testing.T) {
	tests := []struct {
		name string
		i    OxygenSaturationLevel
		want float64
	}{
		{
			name: "Given a OxygenSaturationLevel, OxygenSaturationLevel.Float64() should return its value cast to float64",
			i:    Normal_OxygenSaturationLevel,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.i.Float64())
		})
	}
}

func TestOxygenSaturationLevel_String(t *testing.T) {
	tests := []struct {
		name string
		i    OxygenSaturationLevel
		want string
	}{
		{
			name: "If OxygenSaturationLevel is Normal_OxygenSaturationLevel, OxygenSaturationLevel.String() should return 'normal''",
			i:    Normal_OxygenSaturationLevel,
			want: "normal",
		},
		{
			name: "If OxygenSaturationLevel is Hypoxic_OxygenSaturationLevel, OxygenSaturationLevel.String() should return 'hypoxic''",
			i:    Hypoxic_OxygenSaturationLevel,
			want: "hypoxic",
		},
		{
			name: "If OxygenSaturationLevel is SeverelyHypoxic_OxygenSaturationLevel, OxygenSaturationLevel.String() should return 'severely-hypoxic''",
			i:    SeverelyHypoxic_OxygenSaturationLevel,
			want: "severely-hypoxic",
		},
		{
			name: "If OxygenSaturationLevel is Undefined_OxygenSaturationLevel, OxygenSaturationLevel.String() should return 'undefined''",
			i:    Undefined_OxygenSaturationLevel,
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

func TestNewOxygenSaturationLevelFromString(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want OxygenSaturationLevel
	}{
		{
			name: "Given a string with value 'normal', Normal_OxygenSaturationLevel should be returned",
			args: args{
				s: "normal",
			},
			want: Normal_OxygenSaturationLevel,
		},
		{
			name: "Given a string with value 'hypoxic', Hypoxic_OxygenSaturationLevel should be returned",
			args: args{
				s: "hypoxic",
			},
			want: Hypoxic_OxygenSaturationLevel,
		},
		{
			name: "Given a string with value 'severely-hypoxic', SeverelyHypoxic_OxygenSaturationLevel should be returned",
			args: args{
				s: "severely-hypoxic",
			},
			want: SeverelyHypoxic_OxygenSaturationLevel,
		},
		{
			name: "Given a string unrelated to any OxygenSaturationLevel levels, Undefined_OxygenSaturationLevel should be returned",
			args: args{
				s: "banana",
			},
			want: Undefined_OxygenSaturationLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, NewOxygenSaturationLevelFromString(tt.args.s))
			})
		})
	}
}
