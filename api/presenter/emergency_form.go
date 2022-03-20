package presenter

import "flavioltonon/hmv/domain/valueobject"

type EmergencyForm struct {
	Headache         HeadacheEmergencyFormSession         `json:"headache"`
	Breathing        BreathingEmergencyFormSession        `json:"breathing"`
	ChestPain        ChestPainEmergencyFormSession        `json:"chest_pain"`
	AbdominalPain    AbdominalPainEmergencyFormSession    `json:"abdominal_pain"`
	Backache         BackacheEmergencyFormSession         `json:"backache"`
	BodyTemperature  BodyTemperatureEmergencyFormSession  `json:"body_temperature"`
	BloodPressure    BloodPressureEmergencyFormSession    `json:"blood_pressure"`
	OxygenSaturation OxygenSaturationEmergencyFormSession `json:"oxygen_saturation"`
	Priority         string                               `json:"priority"`
}

func NewEmergencyForm(f valueobject.EmergencyForm) *EmergencyForm {
	return &EmergencyForm{
		Headache:         NewHeadacheEmergencyFormSession(f.Headache),
		Breathing:        NewBreathingEmergencyFormSession(f.Breathing),
		ChestPain:        NewChestPainEmergencyFormSession(f.ChestPain),
		AbdominalPain:    NewAbdominalPainEmergencyFormSession(f.AbdominalPain),
		Backache:         NewBackacheEmergencyFormSession(f.Backache),
		BodyTemperature:  NewBodyTemperatureEmergencyFormSession(f.BodyTemperature),
		BloodPressure:    NewBloodPressureEmergencyFormSession(f.BloodPressure),
		OxygenSaturation: NewOxygenSaturationEmergencyFormSession(f.OxygenSaturation),
		Priority:         f.Priority().String(),
	}
}

type HeadacheEmergencyFormSession struct {
	Has       bool   `json:"has"`
	Intensity string `json:"intensity,omitempty"`
}

func NewHeadacheEmergencyFormSession(session valueobject.HeadacheEmergencyFormSession) HeadacheEmergencyFormSession {
	return HeadacheEmergencyFormSession{Has: session.IsSet(), Intensity: session.HeadacheIntensity.String()}
}

type BreathingEmergencyFormSession struct {
	Has bool `json:"has"`
}

func NewBreathingEmergencyFormSession(session valueobject.BreathingEmergencyFormSession) BreathingEmergencyFormSession {
	return BreathingEmergencyFormSession{Has: session.IsSet()}
}

type ChestPainEmergencyFormSession struct {
	Has             bool   `json:"has"`
	Characteristics string `json:"characteristics,omitempty"`
}

func NewChestPainEmergencyFormSession(session valueobject.ChestPainEmergencyFormSession) ChestPainEmergencyFormSession {
	return ChestPainEmergencyFormSession{Has: session.IsSet(), Characteristics: session.ChestPainCharacteristics.String()}
}

type AbdominalPainEmergencyFormSession struct {
	Has       bool   `json:"has"`
	Intensity string `json:"intensity,omitempty"`
}

func NewAbdominalPainEmergencyFormSession(session valueobject.AbdominalPainEmergencyFormSession) AbdominalPainEmergencyFormSession {
	return AbdominalPainEmergencyFormSession{Has: session.IsSet(), Intensity: session.AbdominalPainIntensity.String()}
}

type BackacheEmergencyFormSession struct {
	Has bool `json:"has"`
}

func NewBackacheEmergencyFormSession(session valueobject.BackacheEmergencyFormSession) BackacheEmergencyFormSession {
	return BackacheEmergencyFormSession{Has: session.IsSet()}
}

type BodyTemperatureEmergencyFormSession struct {
	Temperature *float64 `json:"temperature,omitempty"`
}

func NewBodyTemperatureEmergencyFormSession(session valueobject.BodyTemperatureEmergencyFormSession) BodyTemperatureEmergencyFormSession {
	return BodyTemperatureEmergencyFormSession{Temperature: session.BodyTemperature}
}

type BloodPressureEmergencyFormSession struct {
	Systolic  *float64 `json:"systolic,omitempty"`
	Diastolic *float64 `json:"diastolic,omitempty"`
}

func NewBloodPressureEmergencyFormSession(session valueobject.BloodPressureEmergencyFormSession) BloodPressureEmergencyFormSession {
	return BloodPressureEmergencyFormSession{Systolic: session.SystolicBloodPressure, Diastolic: session.DiastolicBloodPressure}
}

type OxygenSaturationEmergencyFormSession struct {
	Saturation *float64 `json:"saturation,omitempty"`
}

func NewOxygenSaturationEmergencyFormSession(session valueobject.OxygenSaturationEmergencyFormSession) OxygenSaturationEmergencyFormSession {
	return OxygenSaturationEmergencyFormSession{Saturation: session.OxygenSaturation}
}
