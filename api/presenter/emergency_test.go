package presenter

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEmergency(t *testing.T) {
	type args struct {
		e *entity.Emergency
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
		want *Emergency
	}{
		{
			name: "Given a Emergency, a valid presentation should be returned",
			args: args{
				e: &entity.Emergency{
					ID:         "foo",
					PacientID:  "bar",
					Form:       form,
					CreatedAt:  time.Date(2022, time.January, 25, 0, 0, 0, 0, time.UTC),
					UpdatedAt:  time.Date(2022, time.January, 25, 0, 0, 0, 0, time.UTC),
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Triage_EmergencyStatus,
				},
			},
			want: &Emergency{
				ID:        "foo",
				PacientID: "bar",
				Form: &EmergencyForm{
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
				CreatedAt: "25/01/2022 - 00:00:00h",
				UpdatedAt: "25/01/2022 - 00:00:00h",
				Status:    valueobject.Triage_EmergencyStatus.String(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewEmergency(tt.args.e))
		})
	}
}

func TestNewEmergencies(t *testing.T) {
	type args struct {
		es []*entity.Emergency
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
		want Emergencies
	}{
		{
			name: "Given a Emergency, a valid presentation should be returned",
			args: args{
				es: []*entity.Emergency{
					{
						ID:         "foo",
						PacientID:  "bar",
						Form:       form,
						CreatedAt:  time.Date(2022, time.January, 25, 0, 0, 0, 0, time.UTC),
						UpdatedAt:  time.Date(2022, time.January, 25, 0, 0, 0, 0, time.UTC),
						StatusFlow: valueobject.DefaultEmergencyStatusFlow,
						Status:     valueobject.Triage_EmergencyStatus,
					},
				},
			},
			want: Emergencies{
				{
					ID:        "foo",
					PacientID: "bar",
					Form: &EmergencyForm{
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
					CreatedAt: "25/01/2022 - 00:00:00h",
					UpdatedAt: "25/01/2022 - 00:00:00h",
					Status:    valueobject.Triage_EmergencyStatus.String(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, NewEmergencies(tt.args.es))
		})
	}
}
