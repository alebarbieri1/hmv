package presenter

import (
	"flavioltonon/hmv/domain/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHeadacheEmergencyFormSession(t *testing.T) {
	type args struct {
		session valueobject.HeadacheEmergencyFormSession
	}

	_true := true

	tests := []struct {
		name string
		args args
		want *HeadacheEmergencyFormSession
	}{
		{
			name: "Given a HeadacheEmergencyFormSession, a valid presentation should be returned",
			args: args{
				session: valueobject.HeadacheEmergencyFormSession{
					Has:       &_true,
					Intensity: valueobject.VeryHigh_HeadacheIntensity,
				},
			},
			want: &HeadacheEmergencyFormSession{
				Has:       &_true,
				Intensity: valueobject.VeryHigh_HeadacheIntensity.String(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewHeadacheEmergencyFormSession(tt.args.session))
		})
	}
}

func TestNewBreathingDifficultiesEmergencyFormSession(t *testing.T) {
	type args struct {
		session valueobject.BreathingDifficultiesEmergencyFormSession
	}

	_false := false

	tests := []struct {
		name string
		args args
		want *BreathingDifficultiesEmergencyFormSession
	}{
		{
			name: "Given a BreathingDifficultiesEmergencyFormSession, a valid presentation should be returned",
			args: args{
				session: valueobject.BreathingDifficultiesEmergencyFormSession{
					Has: &_false,
				},
			},
			want: &BreathingDifficultiesEmergencyFormSession{
				Has: &_false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewBreathingDifficultiesEmergencyFormSession(tt.args.session))
		})
	}
}

func TestNewChestPainEmergencyFormSession(t *testing.T) {
	type args struct {
		session valueobject.ChestPainEmergencyFormSession
	}

	_false := false

	tests := []struct {
		name string
		args args
		want *ChestPainEmergencyFormSession
	}{
		{
			name: "Given a ChestPainEmergencyFormSession, a valid presentation should be returned",
			args: args{
				session: valueobject.ChestPainEmergencyFormSession{
					Has:             &_false,
					Characteristics: valueobject.Undefined_ChestPainCharacteristics,
				},
			},
			want: &ChestPainEmergencyFormSession{
				Has:             &_false,
				Characteristics: valueobject.Undefined_ChestPainCharacteristics.String(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewChestPainEmergencyFormSession(tt.args.session))
		})
	}
}

func TestNewAbdominalPainEmergencyFormSession(t *testing.T) {
	type args struct {
		session valueobject.AbdominalPainEmergencyFormSession
	}

	_false := false

	tests := []struct {
		name string
		args args
		want *AbdominalPainEmergencyFormSession
	}{
		{
			name: "Given a AbdominalPainEmergencyFormSession, a valid presentation should be returned",
			args: args{
				session: valueobject.AbdominalPainEmergencyFormSession{
					Has:       &_false,
					Intensity: valueobject.Undefined_AbdominalPainIntensity,
				},
			},
			want: &AbdominalPainEmergencyFormSession{
				Has:       &_false,
				Intensity: valueobject.Undefined_AbdominalPainIntensity.String(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewAbdominalPainEmergencyFormSession(tt.args.session))
		})
	}
}

func TestNewBackacheEmergencyFormSession(t *testing.T) {
	type args struct {
		session valueobject.BackacheEmergencyFormSession
	}

	_true := true

	tests := []struct {
		name string
		args args
		want *BackacheEmergencyFormSession
	}{
		{
			name: "Given a BackacheEmergencyFormSession, a valid presentation should be returned",
			args: args{
				session: valueobject.BackacheEmergencyFormSession{
					Has: &_true,
				},
			},
			want: &BackacheEmergencyFormSession{
				Has: &_true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewBackacheEmergencyFormSession(tt.args.session))
		})
	}
}

func TestNewBodyTemperatureEmergencyFormSession(t *testing.T) {
	type args struct {
		session valueobject.BodyTemperatureEmergencyFormSession
	}

	_celsiusDegrees := 36.0

	tests := []struct {
		name string
		args args
		want *BodyTemperatureEmergencyFormSession
	}{
		{
			name: "Given a BodyTemperatureEmergencyFormSession, a valid presentation should be returned",
			args: args{
				session: valueobject.BodyTemperatureEmergencyFormSession{
					CelsiusDegrees: &_celsiusDegrees,
				},
			},
			want: &BodyTemperatureEmergencyFormSession{
				CelsiusDegrees: &_celsiusDegrees,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewBodyTemperatureEmergencyFormSession(tt.args.session))
		})
	}
}

func TestNewBloodPressureEmergencyFormSession(t *testing.T) {
	type args struct {
		session valueobject.BloodPressureEmergencyFormSession
	}

	var (
		_systolic  = 120.0
		_diastolic = 80.0
	)

	tests := []struct {
		name string
		args args
		want *BloodPressureEmergencyFormSession
	}{
		{
			name: "Given a BloodPressureEmergencyFormSession, a valid presentation should be returned",
			args: args{
				session: valueobject.BloodPressureEmergencyFormSession{
					Systolic:  &_systolic,
					Diastolic: &_diastolic,
				},
			},
			want: &BloodPressureEmergencyFormSession{
				Systolic:  &_systolic,
				Diastolic: &_diastolic,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewBloodPressureEmergencyFormSession(tt.args.session))
		})
	}
}

func TestNewOxygenSaturationEmergencyFormSession(t *testing.T) {
	type args struct {
		session valueobject.OxygenSaturationEmergencyFormSession
	}

	_value := 0.97

	tests := []struct {
		name string
		args args
		want *OxygenSaturationEmergencyFormSession
	}{
		{
			name: "Given a OxygenSaturationEmergencyFormSession, a valid presentation should be returned",
			args: args{
				session: valueobject.OxygenSaturationEmergencyFormSession{
					Value: &_value,
				},
			},
			want: &OxygenSaturationEmergencyFormSession{
				Value: &_value,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewOxygenSaturationEmergencyFormSession(tt.args.session))
		})
	}
}

func TestNewEmergencyForm(t *testing.T) {
	type args struct {
		f valueobject.EmergencyForm
	}

	var (
		_true  = true
		_false = false

		_celsiusDegrees = 36.0

		_systolic  = 120.0
		_diastolic = 80.0

		_value = 0.97
	)

	form := valueobject.EmergencyForm{
		Headache: valueobject.HeadacheEmergencyFormSession{
			Has:       &_true,
			Intensity: valueobject.VeryHigh_HeadacheIntensity,
		},
		BreathingDifficulties: valueobject.BreathingDifficultiesEmergencyFormSession{
			Has: &_false,
		},
		ChestPain: valueobject.ChestPainEmergencyFormSession{
			Has:             &_false,
			Characteristics: valueobject.Undefined_ChestPainCharacteristics,
		},
		AbdominalPain: valueobject.AbdominalPainEmergencyFormSession{
			Has:       &_false,
			Intensity: valueobject.Undefined_AbdominalPainIntensity,
		},
		Backache: valueobject.BackacheEmergencyFormSession{
			Has: &_true,
		},
		BodyTemperature: valueobject.BodyTemperatureEmergencyFormSession{
			CelsiusDegrees: &_celsiusDegrees,
		},
		BloodPressure: valueobject.BloodPressureEmergencyFormSession{
			Systolic:  &_systolic,
			Diastolic: &_diastolic,
		},
		OxygenSaturation: valueobject.OxygenSaturationEmergencyFormSession{
			Value: &_value,
		},
	}

	tests := []struct {
		name string
		args args
		want *EmergencyForm
	}{
		{
			name: "Given a EmergencyForm, a valid presentation should be returned",
			args: args{
				f: form,
			},
			want: &EmergencyForm{
				Headache: &HeadacheEmergencyFormSession{
					Has:       &_true,
					Intensity: valueobject.VeryHigh_HeadacheIntensity.String(),
				},
				BreathingDifficulties: &BreathingDifficultiesEmergencyFormSession{
					Has: &_false,
				},
				ChestPain: &ChestPainEmergencyFormSession{
					Has:             &_false,
					Characteristics: valueobject.Undefined_ChestPainCharacteristics.String(),
				},
				AbdominalPain: &AbdominalPainEmergencyFormSession{
					Has:       &_false,
					Intensity: valueobject.Undefined_AbdominalPainIntensity.String(),
				},
				Backache: &BackacheEmergencyFormSession{
					Has: &_true,
				},
				BodyTemperature: &BodyTemperatureEmergencyFormSession{
					CelsiusDegrees: &_celsiusDegrees,
				},
				BloodPressure: &BloodPressureEmergencyFormSession{
					Systolic:  &_systolic,
					Diastolic: &_diastolic,
				},
				OxygenSaturation: &OxygenSaturationEmergencyFormSession{
					Value: &_value,
				},
				Priority: form.Priority().String(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewEmergencyForm(tt.args.f))
		})
	}
}
