package presenter

import "flavioltonon/hmv/domain/valueobject"

type EmergencyForm struct {
	Headache              *HeadacheEmergencyFormSession              `json:"headache,omitempty"`
	BreathingDifficulties *BreathingDifficultiesEmergencyFormSession `json:"breathing_difficulties,omitempty"`
	ChestPain             *ChestPainEmergencyFormSession             `json:"chest_pain,omitempty"`
	AbdominalPain         *AbdominalPainEmergencyFormSession         `json:"abdominal_pain,omitempty"`
	Backache              *BackacheEmergencyFormSession              `json:"backache,omitempty"`
	BodyTemperature       *BodyTemperatureEmergencyFormSession       `json:"body_temperature,omitempty"`
	BloodPressure         *BloodPressureEmergencyFormSession         `json:"blood_pressure,omitempty"`
	OxygenSaturation      *OxygenSaturationEmergencyFormSession      `json:"oxygen_saturation,omitempty"`
	Priority              string                                     `json:"priority,omitempty"`
}

func NewEmergencyForm(f valueobject.EmergencyForm) *EmergencyForm {
	return &EmergencyForm{
		Headache:              NewHeadacheEmergencyFormSession(f.Headache),
		BreathingDifficulties: NewBreathingDifficultiesEmergencyFormSession(f.BreathingDifficulties),
		ChestPain:             NewChestPainEmergencyFormSession(f.ChestPain),
		AbdominalPain:         NewAbdominalPainEmergencyFormSession(f.AbdominalPain),
		Backache:              NewBackacheEmergencyFormSession(f.Backache),
		BodyTemperature:       NewBodyTemperatureEmergencyFormSession(f.BodyTemperature),
		BloodPressure:         NewBloodPressureEmergencyFormSession(f.BloodPressure),
		OxygenSaturation:      NewOxygenSaturationEmergencyFormSession(f.OxygenSaturation),
		Priority:              f.Priority().String(),
	}
}

type HeadacheEmergencyFormSession struct {
	Has       *bool  `json:"has"`
	Intensity string `json:"intensity"`
}

func NewHeadacheEmergencyFormSession(session valueobject.HeadacheEmergencyFormSession) *HeadacheEmergencyFormSession {
	return &HeadacheEmergencyFormSession{Has: session.Has, Intensity: session.Intensity.String()}
}

type BreathingDifficultiesEmergencyFormSession struct {
	Has *bool `json:"has"`
}

func NewBreathingDifficultiesEmergencyFormSession(session valueobject.BreathingDifficultiesEmergencyFormSession) *BreathingDifficultiesEmergencyFormSession {
	return &BreathingDifficultiesEmergencyFormSession{Has: session.Has}
}

type ChestPainEmergencyFormSession struct {
	Has             *bool  `json:"has"`
	Characteristics string `json:"characteristics"`
}

func NewChestPainEmergencyFormSession(session valueobject.ChestPainEmergencyFormSession) *ChestPainEmergencyFormSession {
	return &ChestPainEmergencyFormSession{Has: session.Has, Characteristics: session.Characteristics.String()}
}

type AbdominalPainEmergencyFormSession struct {
	Has       *bool  `json:"has"`
	Intensity string `json:"intensity"`
}

func NewAbdominalPainEmergencyFormSession(session valueobject.AbdominalPainEmergencyFormSession) *AbdominalPainEmergencyFormSession {
	return &AbdominalPainEmergencyFormSession{Has: session.Has, Intensity: session.Intensity.String()}
}

type BackacheEmergencyFormSession struct {
	Has *bool `json:"has"`
}

func NewBackacheEmergencyFormSession(session valueobject.BackacheEmergencyFormSession) *BackacheEmergencyFormSession {
	return &BackacheEmergencyFormSession{Has: session.Has}
}

type BodyTemperatureEmergencyFormSession struct {
	CelsiusDegrees *float64 `json:"celsius_degrees"`
}

func NewBodyTemperatureEmergencyFormSession(session valueobject.BodyTemperatureEmergencyFormSession) *BodyTemperatureEmergencyFormSession {
	return &BodyTemperatureEmergencyFormSession{CelsiusDegrees: session.CelsiusDegrees}
}

type BloodPressureEmergencyFormSession struct {
	Systolic  *float64 `json:"systolic"`
	Diastolic *float64 `json:"diastolic"`
}

func NewBloodPressureEmergencyFormSession(session valueobject.BloodPressureEmergencyFormSession) *BloodPressureEmergencyFormSession {
	return &BloodPressureEmergencyFormSession{Systolic: session.Systolic, Diastolic: session.Diastolic}
}

type OxygenSaturationEmergencyFormSession struct {
	Value *float64 `json:"value"`
}

func NewOxygenSaturationEmergencyFormSession(session valueobject.OxygenSaturationEmergencyFormSession) *OxygenSaturationEmergencyFormSession {
	return &OxygenSaturationEmergencyFormSession{Value: session.Value}
}
