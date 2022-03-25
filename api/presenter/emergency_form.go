package presenter

import "flavioltonon/hmv/domain/valueobject"

// EmergencyForm is a valueobject.EmergencyForm presenter
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

// NewEmergencyForm returns a presentation for a EmergencyForm
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

// HeadacheEmergencyFormSession is a valueobject.HeadacheEmergencyFormSession presenter
type HeadacheEmergencyFormSession struct {
	Has       *bool  `json:"has"`
	Intensity string `json:"intensity"`
}

// NewHeadacheEmergencyFormSession returns a presentation for a HeadacheEmergencyFormSession
func NewHeadacheEmergencyFormSession(session valueobject.HeadacheEmergencyFormSession) *HeadacheEmergencyFormSession {
	return &HeadacheEmergencyFormSession{Has: session.Has, Intensity: session.Intensity.String()}
}

// BreathingDifficultiesEmergencyFormSession is a valueobject.BreathingDifficultiesEmergencyFormSession presenter
type BreathingDifficultiesEmergencyFormSession struct {
	Has *bool `json:"has"`
}

// NewBreathingDifficultiesEmergencyFormSession returns a presentation for a BreathingDifficultiesEmergencyFormSession
func NewBreathingDifficultiesEmergencyFormSession(session valueobject.BreathingDifficultiesEmergencyFormSession) *BreathingDifficultiesEmergencyFormSession {
	return &BreathingDifficultiesEmergencyFormSession{Has: session.Has}
}

// ChestPainEmergencyFormSession is a valueobject.ChestPainEmergencyFormSession presenter
type ChestPainEmergencyFormSession struct {
	Has             *bool  `json:"has"`
	Characteristics string `json:"characteristics"`
}

// NewChestPainEmergencyFormSession returns a presentation for a ChestPainEmergencyFormSession
func NewChestPainEmergencyFormSession(session valueobject.ChestPainEmergencyFormSession) *ChestPainEmergencyFormSession {
	return &ChestPainEmergencyFormSession{Has: session.Has, Characteristics: session.Characteristics.String()}
}

// AbdominalPainEmergencyFormSession is a valueobject.AbdominalPainEmergencyFormSession presenter
type AbdominalPainEmergencyFormSession struct {
	Has       *bool  `json:"has"`
	Intensity string `json:"intensity"`
}

// NewAbdominalPainEmergencyFormSession returns a presentation for a AbdominalPainEmergencyFormSession
func NewAbdominalPainEmergencyFormSession(session valueobject.AbdominalPainEmergencyFormSession) *AbdominalPainEmergencyFormSession {
	return &AbdominalPainEmergencyFormSession{Has: session.Has, Intensity: session.Intensity.String()}
}

// BackacheEmergencyFormSession is a valueobject.BackacheEmergencyFormSession presenter
type BackacheEmergencyFormSession struct {
	Has *bool `json:"has"`
}

// NewBackacheEmergencyFormSession returns a presentation for a BackacheEmergencyFormSession
func NewBackacheEmergencyFormSession(session valueobject.BackacheEmergencyFormSession) *BackacheEmergencyFormSession {
	return &BackacheEmergencyFormSession{Has: session.Has}
}

// BodyTemperatureEmergencyFormSession is a valueobject.BodyTemperatureEmergencyFormSession presenter
type BodyTemperatureEmergencyFormSession struct {
	CelsiusDegrees *float64 `json:"celsius_degrees"`
}

// NewBodyTemperatureEmergencyFormSession returns a presentation for a BodyTemperatureEmergencyFormSession
func NewBodyTemperatureEmergencyFormSession(session valueobject.BodyTemperatureEmergencyFormSession) *BodyTemperatureEmergencyFormSession {
	return &BodyTemperatureEmergencyFormSession{CelsiusDegrees: session.CelsiusDegrees}
}

// BloodPressureEmergencyFormSession is a valueobject.BloodPressureEmergencyFormSession presenter
type BloodPressureEmergencyFormSession struct {
	Systolic  *float64 `json:"systolic"`
	Diastolic *float64 `json:"diastolic"`
}

// NewBloodPressureEmergencyFormSession returns a presentation for a BloodPressureEmergencyFormSession
func NewBloodPressureEmergencyFormSession(session valueobject.BloodPressureEmergencyFormSession) *BloodPressureEmergencyFormSession {
	return &BloodPressureEmergencyFormSession{Systolic: session.Systolic, Diastolic: session.Diastolic}
}

// OxygenSaturationEmergencyFormSession is a valueobject.OxygenSaturationEmergencyFormSession presenter
type OxygenSaturationEmergencyFormSession struct {
	Value *float64 `json:"value"`
}

// NewOxygenSaturationEmergencyFormSession returns a presentation for a OxygenSaturationEmergencyFormSession
func NewOxygenSaturationEmergencyFormSession(session valueobject.OxygenSaturationEmergencyFormSession) *OxygenSaturationEmergencyFormSession {
	return &OxygenSaturationEmergencyFormSession{Value: session.Value}
}
