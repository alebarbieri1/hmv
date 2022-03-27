package pacients

import (
	"encoding/json"
	"errors"
	"net/http"

	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/mux"
)

func (c *Controller) updateHealthData(w http.ResponseWriter, r *http.Request) {
	user, err := entity.NewUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.FailedToAuthenticateUser, err))
		return
	}

	payload, err := c.newUpdateHealthDataPayload(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToValidateRequest, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToValidateRequest, err))
		return
	}

	vars := mux.Vars(r)

	updatedPacient, err := c.usecases.Pacients.UpdateHealthData(user.ID, vars["pacient_id"], payload.toValueObject())
	if errors.Is(err, entity.ErrNotFound) {
		c.drivers.Logger.Info(application.FailedToUpdatePacient, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToUpdatePacient, err))
		return
	}

	if errors.Is(err, application.ErrInvalidUserProfile) {
		c.drivers.Logger.Info(application.FailedToUpdatePacient, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToUpdatePacient, err))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToUpdatePacient, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToUpdatePacient, err))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewPacient(updatedPacient)))
}

type updateHealthDataPayload struct {
	AllergicToMedicationDescription          string `json:"allergic_to_medication_description"`
	ContinuousUseMedicationDescription       string `json:"continuous_use_medication_description"`
	DiagnosedDiseaseDescription              string `json:"diagnosed_disease_description"`
	HistoryOfHeartDiseaseDescription         string `json:"history_of_heart_disease_description"`
	SmokingBehaviorDescription               string `json:"smoking_behavior_description"`
	PhysicalActivityDescription              string `json:"physical_activity_description"`
	AlcoholicBeveragesConsumptionDescription string `json:"alcoholic_beverage_consumption_description"`
}

func (p *updateHealthDataPayload) Validate() error {
	return ozzo.ValidateStruct(p,
		ozzo.Field(&p.AllergicToMedicationDescription, ozzo.Length(0, 256)),
		ozzo.Field(&p.ContinuousUseMedicationDescription, ozzo.Length(0, 256)),
		ozzo.Field(&p.DiagnosedDiseaseDescription, ozzo.Length(0, 256)),
		ozzo.Field(&p.HistoryOfHeartDiseaseDescription, ozzo.Length(0, 256)),
		ozzo.Field(&p.SmokingBehaviorDescription, ozzo.Length(0, 256)),
		ozzo.Field(&p.PhysicalActivityDescription, ozzo.Length(0, 256)),
		ozzo.Field(&p.AlcoholicBeveragesConsumptionDescription, ozzo.Length(0, 256)),
	)
}

func (p *updateHealthDataPayload) toValueObject() valueobject.HealthData {
	return valueobject.HealthData{
		AllergicToMedicationDescription:          p.AllergicToMedicationDescription,
		ContinuousUseMedicationDescription:       p.ContinuousUseMedicationDescription,
		DiagnosedDiseaseDescription:              p.DiagnosedDiseaseDescription,
		HistoryOfHeartDiseaseDescription:         p.HistoryOfHeartDiseaseDescription,
		SmokingBehaviorDescription:               p.SmokingBehaviorDescription,
		PhysicalActivityDescription:              p.PhysicalActivityDescription,
		AlcoholicBeveragesConsumptionDescription: p.AlcoholicBeveragesConsumptionDescription,
	}
}

func (c *Controller) newUpdateHealthDataPayload(r *http.Request) (*updateHealthDataPayload, error) {
	var payload updateHealthDataPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, err
	}

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	return &payload, nil
}
