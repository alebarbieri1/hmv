package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmergencyForm_Validate(t *testing.T) {
	type fields struct {
		Headache         HeadacheEmergencyFormSession
		ChestPain        ChestPainEmergencyFormSession
		AbdominalPain    AbdominalPainEmergencyFormSession
		OxygenSaturation OxygenSaturationEmergencyFormSession
	}

	var (
		_true  = true
		_false = false

		_value         = 0.97
		_negativeValue = -1.0
	)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "If all EmergencyForm sessions are valid, Validate should return no errors",
			fields: fields{
				Headache: HeadacheEmergencyFormSession{
					Has:       &_true,
					Intensity: VeryHigh_HeadacheIntensity,
				},
				ChestPain: ChestPainEmergencyFormSession{
					Has:             &_false,
					Characteristics: Undefined_ChestPainCharacteristics,
				},
				AbdominalPain: AbdominalPainEmergencyFormSession{
					Has:       &_false,
					Intensity: Undefined_AbdominalPainIntensity,
				},
				OxygenSaturation: OxygenSaturationEmergencyFormSession{
					Value: &_value,
				},
			},
			wantErr: false,
		},
		{
			name: "If Headache is not valid, Validate should return an error",
			fields: fields{
				Headache: HeadacheEmergencyFormSession{
					Has:       &_false,
					Intensity: VeryHigh_HeadacheIntensity,
				},
				ChestPain: ChestPainEmergencyFormSession{
					Has:             &_false,
					Characteristics: Undefined_ChestPainCharacteristics,
				},
				AbdominalPain: AbdominalPainEmergencyFormSession{
					Has:       &_false,
					Intensity: Undefined_AbdominalPainIntensity,
				},
				OxygenSaturation: OxygenSaturationEmergencyFormSession{
					Value: &_value,
				},
			},
			wantErr: true,
		},
		{
			name: "If ChestPain is not valid, Validate should return an error",
			fields: fields{
				Headache: HeadacheEmergencyFormSession{
					Has:       &_true,
					Intensity: VeryHigh_HeadacheIntensity,
				},
				ChestPain: ChestPainEmergencyFormSession{
					Has:             &_false,
					Characteristics: RadiatingToTheLeftArm_ChestPainCharacteristics,
				},
				AbdominalPain: AbdominalPainEmergencyFormSession{
					Has:       &_false,
					Intensity: Undefined_AbdominalPainIntensity,
				},
				OxygenSaturation: OxygenSaturationEmergencyFormSession{
					Value: &_value,
				},
			},
			wantErr: true,
		},
		{
			name: "If AbdominalPain is not valid, Validate should return an error",
			fields: fields{
				Headache: HeadacheEmergencyFormSession{
					Has:       &_true,
					Intensity: VeryHigh_HeadacheIntensity,
				},
				ChestPain: ChestPainEmergencyFormSession{
					Has:             &_false,
					Characteristics: Undefined_ChestPainCharacteristics,
				},
				AbdominalPain: AbdominalPainEmergencyFormSession{
					Has:       &_false,
					Intensity: VeryHigh_AbdominalPainIntensity,
				},
				OxygenSaturation: OxygenSaturationEmergencyFormSession{
					Value: &_value,
				},
			},
			wantErr: true,
		},
		{
			name: "If OxygenSaturation is not valid, Validate should return an error",
			fields: fields{
				Headache: HeadacheEmergencyFormSession{
					Has:       &_true,
					Intensity: VeryHigh_HeadacheIntensity,
				},
				ChestPain: ChestPainEmergencyFormSession{
					Has:             &_false,
					Characteristics: Undefined_ChestPainCharacteristics,
				},
				AbdominalPain: AbdominalPainEmergencyFormSession{
					Has:       &_false,
					Intensity: Undefined_AbdominalPainIntensity,
				},
				OxygenSaturation: OxygenSaturationEmergencyFormSession{
					Value: &_negativeValue,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := EmergencyForm{
				Headache:         tt.fields.Headache,
				ChestPain:        tt.fields.ChestPain,
				AbdominalPain:    tt.fields.AbdominalPain,
				OxygenSaturation: tt.fields.OxygenSaturation,
			}

			err := f.Validate()

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestNewEmergencyForm(t *testing.T) {
	type args struct {
		headache              HeadacheEmergencyFormSession
		breathingDifficulties BreathingDifficultiesEmergencyFormSession
		chestPain             ChestPainEmergencyFormSession
		abdominalPain         AbdominalPainEmergencyFormSession
		backache              BackacheEmergencyFormSession
		bodyTemperature       BodyTemperatureEmergencyFormSession
		bloodPressure         BloodPressureEmergencyFormSession
		oxygenSaturation      OxygenSaturationEmergencyFormSession
	}

	var (
		_true  = true
		_false = false

		_celsiusDegrees = 36.0

		_systolic  = 120.0
		_diastolic = 80.0

		_value = 0.97
	)

	tests := []struct {
		name    string
		args    args
		want    EmergencyForm
		wantErr bool
	}{
		{
			name: "Given a set of valid EmergencyFormSession, a new EmergencyForm should be created",
			args: args{
				headache: HeadacheEmergencyFormSession{
					Has:       &_true,
					Intensity: VeryHigh_HeadacheIntensity,
				},
				breathingDifficulties: BreathingDifficultiesEmergencyFormSession{
					Has: &_false,
				},
				chestPain: ChestPainEmergencyFormSession{
					Has:             &_false,
					Characteristics: Undefined_ChestPainCharacteristics,
				},
				abdominalPain: AbdominalPainEmergencyFormSession{
					Has:       &_false,
					Intensity: Undefined_AbdominalPainIntensity,
				},
				backache: BackacheEmergencyFormSession{
					Has: &_true,
				},
				bodyTemperature: BodyTemperatureEmergencyFormSession{
					CelsiusDegrees: &_celsiusDegrees,
				},
				bloodPressure: BloodPressureEmergencyFormSession{
					Systolic:  &_systolic,
					Diastolic: &_diastolic,
				},
				oxygenSaturation: OxygenSaturationEmergencyFormSession{
					Value: &_value,
				},
			},
			want: EmergencyForm{
				Headache: HeadacheEmergencyFormSession{
					Has:       &_true,
					Intensity: VeryHigh_HeadacheIntensity,
				},
				BreathingDifficulties: BreathingDifficultiesEmergencyFormSession{
					Has: &_false,
				},
				ChestPain: ChestPainEmergencyFormSession{
					Has:             &_false,
					Characteristics: Undefined_ChestPainCharacteristics,
				},
				AbdominalPain: AbdominalPainEmergencyFormSession{
					Has:       &_false,
					Intensity: Undefined_AbdominalPainIntensity,
				},
				Backache: BackacheEmergencyFormSession{
					Has: &_true,
				},
				BodyTemperature: BodyTemperatureEmergencyFormSession{
					CelsiusDegrees: &_celsiusDegrees,
				},
				BloodPressure: BloodPressureEmergencyFormSession{
					Systolic:  &_systolic,
					Diastolic: &_diastolic,
				},
				OxygenSaturation: OxygenSaturationEmergencyFormSession{
					Value: &_value,
				},
			},
			wantErr: false,
		},
		{
			name: "Given a set containing invalid any EmergencyFormSession, an error should be returned",
			args: args{
				headache: HeadacheEmergencyFormSession{
					Has:       &_true,
					Intensity: Undefined_HeadacheIntensity,
				},
				breathingDifficulties: BreathingDifficultiesEmergencyFormSession{
					Has: &_false,
				},
				chestPain: ChestPainEmergencyFormSession{
					Has:             &_false,
					Characteristics: Undefined_ChestPainCharacteristics,
				},
				abdominalPain: AbdominalPainEmergencyFormSession{
					Has:       &_false,
					Intensity: Undefined_AbdominalPainIntensity,
				},
				backache: BackacheEmergencyFormSession{
					Has: &_true,
				},
				bodyTemperature: BodyTemperatureEmergencyFormSession{
					CelsiusDegrees: &_celsiusDegrees,
				},
				bloodPressure: BloodPressureEmergencyFormSession{
					Systolic:  &_systolic,
					Diastolic: &_diastolic,
				},
				oxygenSaturation: OxygenSaturationEmergencyFormSession{
					Value: &_value,
				},
			},
			want:    EmergencyForm{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEmergencyForm(
				tt.args.headache,
				tt.args.breathingDifficulties,
				tt.args.chestPain,
				tt.args.abdominalPain,
				tt.args.backache,
				tt.args.bodyTemperature,
				tt.args.bloodPressure,
				tt.args.oxygenSaturation,
			)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEmergencyForm_Priority(t *testing.T) {
	type fields struct {
		Headache              HeadacheEmergencyFormSession
		BreathingDifficulties BreathingDifficultiesEmergencyFormSession
		ChestPain             ChestPainEmergencyFormSession
		AbdominalPain         AbdominalPainEmergencyFormSession
		Backache              BackacheEmergencyFormSession
		BodyTemperature       BodyTemperatureEmergencyFormSession
		BloodPressure         BloodPressureEmergencyFormSession
		OxygenSaturation      OxygenSaturationEmergencyFormSession
	}

	_true := true

	tests := []struct {
		name   string
		fields fields
		want   EmergencyPriority
	}{
		{
			name: "Given an empty EmergencyForm, Undefined_EmergencyPriority should be returned",
			want: Undefined_EmergencyPriority,
		},
		{
			name: "Given an EmergencyForm with very high priority, VeryHigh_EmergencyPriority should be returned",
			fields: fields{
				ChestPain: ChestPainEmergencyFormSession{
					Has:             &_true,
					Characteristics: RadiatingToTheLeftArm_ChestPainCharacteristics,
				},
			},
			want: VeryHigh_EmergencyPriority,
		},
		{
			name: "Given an EmergencyForm with high priority, High_EmergencyPriority should be returned",
			fields: fields{
				AbdominalPain: AbdominalPainEmergencyFormSession{
					Has:       &_true,
					Intensity: VeryHigh_AbdominalPainIntensity,
				},
			},
			want: High_EmergencyPriority,
		},
		{
			name: "Given an EmergencyForm with medium priority, Medium_EmergencyPriority should be returned",
			fields: fields{
				AbdominalPain: AbdominalPainEmergencyFormSession{
					Has:       &_true,
					Intensity: Medium_AbdominalPainIntensity,
				},
			},
			want: Medium_EmergencyPriority,
		},
		{
			name: "Given an EmergencyForm with low priority, Low_EmergencyPriority should be returned",
			fields: fields{
				AbdominalPain: AbdominalPainEmergencyFormSession{
					Has:       &_true,
					Intensity: Low_AbdominalPainIntensity,
				},
			},
			want: Low_EmergencyPriority,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := EmergencyForm{
				Headache:              tt.fields.Headache,
				BreathingDifficulties: tt.fields.BreathingDifficulties,
				ChestPain:             tt.fields.ChestPain,
				AbdominalPain:         tt.fields.AbdominalPain,
				Backache:              tt.fields.Backache,
				BodyTemperature:       tt.fields.BodyTemperature,
				BloodPressure:         tt.fields.BloodPressure,
				OxygenSaturation:      tt.fields.OxygenSaturation,
			}

			assert.Equal(t, tt.want, f.Priority())
		})
	}
}

func TestEmergencyForm_hasVeryHighPriority(t *testing.T) {
	type fields struct {
		Headache              HeadacheEmergencyFormSession
		BreathingDifficulties BreathingDifficultiesEmergencyFormSession
		ChestPain             ChestPainEmergencyFormSession
		AbdominalPain         AbdominalPainEmergencyFormSession
		Backache              BackacheEmergencyFormSession
		BodyTemperature       BodyTemperatureEmergencyFormSession
		BloodPressure         BloodPressureEmergencyFormSession
		OxygenSaturation      OxygenSaturationEmergencyFormSession
	}

	var (
		_true = true

		_systolic  = 104.0
		_diastolic = 73.0

		_value = 0.9
	)

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "If ChestPain.Score() is higher than 0.5, EmergencyForm.hasVeryHighPriority() should return true",
			fields: fields{
				ChestPain: ChestPainEmergencyFormSession{
					Has:             &_true,
					Characteristics: RadiatingToTheLeftArm_ChestPainCharacteristics,
				},
			},
			want: true,
		},
		{
			name: "If BloodPressure.Score() is higher than 0.7, EmergencyForm.hasVeryHighPriority() should return true",
			fields: fields{
				BloodPressure: BloodPressureEmergencyFormSession{
					Systolic:  &_systolic,
					Diastolic: &_diastolic,
				},
			},
			want: true,
		},
		{
			name: "If OxygenSaturation.Score() is higher than 0.5, EmergencyForm.hasVeryHighPriority() should return true",
			fields: fields{
				OxygenSaturation: OxygenSaturationEmergencyFormSession{
					Value: &_value,
				},
			},
			want: true,
		},
		{
			name:   "If none of the requirements are met, EmergencyForm.hasVeryHighPriority() should return false",
			fields: fields{},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := EmergencyForm{
				Headache:              tt.fields.Headache,
				BreathingDifficulties: tt.fields.BreathingDifficulties,
				ChestPain:             tt.fields.ChestPain,
				AbdominalPain:         tt.fields.AbdominalPain,
				Backache:              tt.fields.Backache,
				BodyTemperature:       tt.fields.BodyTemperature,
				BloodPressure:         tt.fields.BloodPressure,
				OxygenSaturation:      tt.fields.OxygenSaturation,
			}

			assert.Equal(t, tt.want, f.hasVeryHighPriority())
		})
	}
}

func TestEmergencyForm_hasHighPriority(t *testing.T) {
	type fields struct {
		Headache              HeadacheEmergencyFormSession
		BreathingDifficulties BreathingDifficultiesEmergencyFormSession
		ChestPain             ChestPainEmergencyFormSession
		AbdominalPain         AbdominalPainEmergencyFormSession
		Backache              BackacheEmergencyFormSession
		BodyTemperature       BodyTemperatureEmergencyFormSession
		BloodPressure         BloodPressureEmergencyFormSession
		OxygenSaturation      OxygenSaturationEmergencyFormSession
	}

	var (
		_true = true

		_celsiusDegrees = 50.0

		_systolic  = 104.0
		_diastolic = 73.0
	)

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "If BloodPressure.Score() is higher than 0.5, EmergencyForm.hasHighPriority() should return true",
			fields: fields{
				BloodPressure: BloodPressureEmergencyFormSession{
					Systolic:  &_systolic,
					Diastolic: &_diastolic,
				},
			},
			want: true,
		},
		{
			name: "If BodyTemperature.Score() is higher than 0.75, EmergencyForm.hasHighPriority() should return true",
			fields: fields{
				BodyTemperature: BodyTemperatureEmergencyFormSession{
					CelsiusDegrees: &_celsiusDegrees,
				},
			},
			want: true,
		},
		{
			name: "If AbdominalPain.Score() is higher than 0.7, EmergencyForm.hasHighPriority() should return true",
			fields: fields{
				AbdominalPain: AbdominalPainEmergencyFormSession{
					Has:       &_true,
					Intensity: VeryHigh_AbdominalPainIntensity,
				},
			},
			want: true,
		},
		{
			name:   "If none of the requirements are met, EmergencyForm.hasHighPriority() should return false",
			fields: fields{},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := EmergencyForm{
				Headache:              tt.fields.Headache,
				BreathingDifficulties: tt.fields.BreathingDifficulties,
				ChestPain:             tt.fields.ChestPain,
				AbdominalPain:         tt.fields.AbdominalPain,
				Backache:              tt.fields.Backache,
				BodyTemperature:       tt.fields.BodyTemperature,
				BloodPressure:         tt.fields.BloodPressure,
				OxygenSaturation:      tt.fields.OxygenSaturation,
			}

			assert.Equal(t, tt.want, f.hasHighPriority())
		})
	}
}

func TestEmergencyForm_hasMediumPriority(t *testing.T) {
	type fields struct {
		Headache              HeadacheEmergencyFormSession
		BreathingDifficulties BreathingDifficultiesEmergencyFormSession
		ChestPain             ChestPainEmergencyFormSession
		AbdominalPain         AbdominalPainEmergencyFormSession
		Backache              BackacheEmergencyFormSession
		BodyTemperature       BodyTemperatureEmergencyFormSession
		BloodPressure         BloodPressureEmergencyFormSession
		OxygenSaturation      OxygenSaturationEmergencyFormSession
	}

	var (
		_true = true

		_celsiusDegrees = 50.0
	)

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "If BodyTemperature.Score() is higher than 0.5, EmergencyForm.hasMediumPriority() should return true",
			fields: fields{
				BodyTemperature: BodyTemperatureEmergencyFormSession{
					CelsiusDegrees: &_celsiusDegrees,
				},
			},
			want: true,
		},
		{
			name: "If BreathingDifficulties.Score() is higher than 0.5, EmergencyForm.hasMediumPriority() should return true",
			fields: fields{
				BreathingDifficulties: BreathingDifficultiesEmergencyFormSession{
					Has: &_true,
				},
			},
			want: true,
		},
		{
			name: "If Headache.Score() is higher than 0.7, EmergencyForm.hasMediumPriority() should return true",
			fields: fields{
				Headache: HeadacheEmergencyFormSession{
					Has:       &_true,
					Intensity: High_HeadacheIntensity,
				},
			},
			want: true,
		},
		{
			name: "If Backache.Score() is higher than 0.5, EmergencyForm.hasMediumPriority() should return true",
			fields: fields{
				Backache: BackacheEmergencyFormSession{
					Has: &_true,
				},
			},
			want: true,
		},
		{
			name:   "If none of the requirements are met, EmergencyForm.hasMediumPriority() should return false",
			fields: fields{},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := EmergencyForm{
				Headache:              tt.fields.Headache,
				BreathingDifficulties: tt.fields.BreathingDifficulties,
				ChestPain:             tt.fields.ChestPain,
				AbdominalPain:         tt.fields.AbdominalPain,
				Backache:              tt.fields.Backache,
				BodyTemperature:       tt.fields.BodyTemperature,
				BloodPressure:         tt.fields.BloodPressure,
				OxygenSaturation:      tt.fields.OxygenSaturation,
			}

			assert.Equal(t, tt.want, f.hasMediumPriority())
		})
	}
}

func TestEmergencyForm_IsEmpty(t *testing.T) {
	type fields struct {
		Headache              HeadacheEmergencyFormSession
		BreathingDifficulties BreathingDifficultiesEmergencyFormSession
		ChestPain             ChestPainEmergencyFormSession
		AbdominalPain         AbdominalPainEmergencyFormSession
		Backache              BackacheEmergencyFormSession
		BodyTemperature       BodyTemperatureEmergencyFormSession
		BloodPressure         BloodPressureEmergencyFormSession
		OxygenSaturation      OxygenSaturationEmergencyFormSession
	}

	_true := true

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "If no EmergencyFormSession are set, EmergencyForm.IsEmpty() should return true",
			fields: fields{},
			want:   true,
		},
		{
			name: "If any EmergencyFormSession are set, EmergencyForm.IsEmpty() should return false",
			fields: fields{
				Headache: HeadacheEmergencyFormSession{
					Has:       &_true,
					Intensity: High_HeadacheIntensity,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := EmergencyForm{
				Headache:              tt.fields.Headache,
				BreathingDifficulties: tt.fields.BreathingDifficulties,
				ChestPain:             tt.fields.ChestPain,
				AbdominalPain:         tt.fields.AbdominalPain,
				Backache:              tt.fields.Backache,
				BodyTemperature:       tt.fields.BodyTemperature,
				BloodPressure:         tt.fields.BloodPressure,
				OxygenSaturation:      tt.fields.OxygenSaturation,
			}

			assert.Equal(t, tt.want, f.IsEmpty())
		})
	}
}

func TestEmergencyForm_IsComplete(t *testing.T) {
	type fields struct {
		Headache              HeadacheEmergencyFormSession
		BreathingDifficulties BreathingDifficultiesEmergencyFormSession
		ChestPain             ChestPainEmergencyFormSession
		AbdominalPain         AbdominalPainEmergencyFormSession
		Backache              BackacheEmergencyFormSession
		BodyTemperature       BodyTemperatureEmergencyFormSession
		BloodPressure         BloodPressureEmergencyFormSession
		OxygenSaturation      OxygenSaturationEmergencyFormSession
	}

	var (
		_true  = true
		_false = false

		_celsiusDegrees = 36.0

		_systolic  = 120.0
		_diastolic = 80.0

		_value = 0.97
	)

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "If all EmergencyFormSession are set, EmergencyForm.IsComplete() should return true",
			fields: fields{
				Headache: HeadacheEmergencyFormSession{
					Has:       &_true,
					Intensity: VeryHigh_HeadacheIntensity,
				},
				BreathingDifficulties: BreathingDifficultiesEmergencyFormSession{
					Has: &_false,
				},
				ChestPain: ChestPainEmergencyFormSession{
					Has:             &_false,
					Characteristics: Undefined_ChestPainCharacteristics,
				},
				AbdominalPain: AbdominalPainEmergencyFormSession{
					Has:       &_false,
					Intensity: Undefined_AbdominalPainIntensity,
				},
				Backache: BackacheEmergencyFormSession{
					Has: &_true,
				},
				BodyTemperature: BodyTemperatureEmergencyFormSession{
					CelsiusDegrees: &_celsiusDegrees,
				},
				BloodPressure: BloodPressureEmergencyFormSession{
					Systolic:  &_systolic,
					Diastolic: &_diastolic,
				},
				OxygenSaturation: OxygenSaturationEmergencyFormSession{
					Value: &_value,
				},
			},
			want: true,
		},
		{
			name: "If any EmergencyFormSession are not set, EmergencyForm.IsComplete() should return false",
			fields: fields{
				Headache: HeadacheEmergencyFormSession{
					Has:       &_true,
					Intensity: High_HeadacheIntensity,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := EmergencyForm{
				Headache:              tt.fields.Headache,
				BreathingDifficulties: tt.fields.BreathingDifficulties,
				ChestPain:             tt.fields.ChestPain,
				AbdominalPain:         tt.fields.AbdominalPain,
				Backache:              tt.fields.Backache,
				BodyTemperature:       tt.fields.BodyTemperature,
				BloodPressure:         tt.fields.BloodPressure,
				OxygenSaturation:      tt.fields.OxygenSaturation,
			}

			assert.Equal(t, tt.want, f.IsComplete())
		})
	}
}
