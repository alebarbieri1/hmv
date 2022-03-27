package presenter

import "flavioltonon/hmv/domain/valueobject"

type HealthData struct {
	AllergicToMedicationDescription          string `json:"allergic_to_medication_description"`
	ContinuousUseMedicationDescription       string `json:"continuous_use_medication_description"`
	DiagnosedDiseaseDescription              string `json:"diagnosed_disease_description"`
	HistoryOfHeartDiseaseDescription         string `json:"history_of_heart_disease_description"`
	SmokingBehaviorDescription               string `json:"smoking_behavior_description"`
	PhysicalActivityDescription              string `json:"physical_activity_description"`
	AlcoholicBeveragesConsumptionDescription string `json:"alcoholic_beverage_consumption_description"`
}

func NewHealthData(data valueobject.HealthData) *HealthData {
	return &HealthData{
		AllergicToMedicationDescription:          data.AllergicToMedicationDescription,
		ContinuousUseMedicationDescription:       data.ContinuousUseMedicationDescription,
		DiagnosedDiseaseDescription:              data.DiagnosedDiseaseDescription,
		HistoryOfHeartDiseaseDescription:         data.HistoryOfHeartDiseaseDescription,
		SmokingBehaviorDescription:               data.SmokingBehaviorDescription,
		PhysicalActivityDescription:              data.PhysicalActivityDescription,
		AlcoholicBeveragesConsumptionDescription: data.AlcoholicBeveragesConsumptionDescription,
	}
}
