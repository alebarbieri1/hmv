package valueobject

import (
	"reflect"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type EmergencyForm struct {
	Headache              HeadacheEmergencyFormSession
	BreathingDifficulties BreathingDifficultiesEmergencyFormSession
	ChestPain             ChestPainEmergencyFormSession
	AbdominalPain         AbdominalPainEmergencyFormSession
	Backache              BackacheEmergencyFormSession
	BodyTemperature       BodyTemperatureEmergencyFormSession
	BloodPressure         BloodPressureEmergencyFormSession
	OxygenSaturation      OxygenSaturationEmergencyFormSession
}

type EmergencyFormSession interface {
	IsSet() bool
	Score() float64
}

func NewEmergencyForm(
	headache HeadacheEmergencyFormSession,
	breathingDifficulties BreathingDifficultiesEmergencyFormSession,
	chestPain ChestPainEmergencyFormSession,
	abdominalPain AbdominalPainEmergencyFormSession,
	backache BackacheEmergencyFormSession,
	bodyTemperature BodyTemperatureEmergencyFormSession,
	bloodPressure BloodPressureEmergencyFormSession,
	oxygenSaturation OxygenSaturationEmergencyFormSession,
) (EmergencyForm, error) {
	form := EmergencyForm{
		Headache:              headache,
		BreathingDifficulties: breathingDifficulties,
		ChestPain:             chestPain,
		AbdominalPain:         abdominalPain,
		Backache:              backache,
		BodyTemperature:       bodyTemperature,
		BloodPressure:         bloodPressure,
		OxygenSaturation:      oxygenSaturation,
	}

	if err := form.Validate(); err != nil {
		return EmergencyForm{}, err
	}

	return form, nil
}

func (f EmergencyForm) Validate() error {
	return ozzo.ValidateStruct(&f,
		ozzo.Field(&f.Headache),
		ozzo.Field(&f.BreathingDifficulties),
		ozzo.Field(&f.ChestPain),
		ozzo.Field(&f.AbdominalPain),
		ozzo.Field(&f.Backache),
		ozzo.Field(&f.BodyTemperature),
		ozzo.Field(&f.BloodPressure),
		ozzo.Field(&f.OxygenSaturation),
	)
}

func (f EmergencyForm) Priority() EmergencyPriority {
	if f.IsEmpty() {
		return Undefined_EmergencyPriority
	}

	switch {
	case f.hasVeryHighPriority():
		return VeryHigh_EmergencyPriority
	case f.hasHighPriority():
		return High_EmergencyPriority
	case f.hasMediumPriority():
		return Medium_EmergencyPriority
	}

	return Low_EmergencyPriority
}

func (f EmergencyForm) hasVeryHighPriority() bool {
	return f.ChestPain.Score() > 0.5 ||
		f.BloodPressure.Score() > 0.7 ||
		f.OxygenSaturation.Score() > 0.5
}

func (f EmergencyForm) hasHighPriority() bool {
	return f.BloodPressure.Score() > 0.5 ||
		f.BodyTemperature.Score() > 0.75 ||
		f.AbdominalPain.Score() > 0.7
}

func (f EmergencyForm) hasMediumPriority() bool {
	return f.BodyTemperature.Score() > 0.5 ||
		f.AbdominalPain.Score() > 0.5 ||
		f.BreathingDifficulties.Score() > 0.5 ||
		f.Headache.Score() > 0.7 ||
		f.Backache.Score() > 0.5
}

func (f EmergencyForm) IsEmpty() bool {
	v := reflect.ValueOf(f)

	for i := 0; i < v.NumField(); i++ {
		if field, implements := v.Field(i).Interface().(EmergencyFormSession); implements && field.IsSet() {
			return false
		}
	}

	return true
}

func (f EmergencyForm) IsComplete() bool {
	v := reflect.ValueOf(f)

	for i := 0; i < v.NumField(); i++ {
		if field, implements := v.Field(i).Interface().(EmergencyFormSession); !implements || !field.IsSet() {
			return false
		}
	}

	return true
}
