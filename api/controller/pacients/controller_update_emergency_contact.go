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

func (c *Controller) updateEmergencyContact(w http.ResponseWriter, r *http.Request) {
	user, err := entity.NewUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.FailedToAuthenticateUser, err))
		return
	}

	payload, err := c.newUpdateEmergencyContactPayload(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToValidateRequest, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToValidateRequest, err))
		return
	}

	vars := mux.Vars(r)

	updatedPacient, err := c.usecases.Pacients.UpdateEmergencyContact(user.ID, vars["pacient_id"], payload.toValueObject())
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

type updateEmergencyContactPayload struct {
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
}

func (p *updateEmergencyContactPayload) Validate() error {
	return ozzo.ValidateStruct(p,
		ozzo.Field(&p.Name, ozzo.Required),
		ozzo.Field(&p.MobileNumber, ozzo.Required),
	)
}

func (p *updateEmergencyContactPayload) toValueObject() valueobject.EmergencyContact {
	return valueobject.EmergencyContact{
		Name:         p.Name,
		MobileNumber: p.MobileNumber,
	}
}

func (c *Controller) newUpdateEmergencyContactPayload(r *http.Request) (*updateEmergencyContactPayload, error) {
	var payload updateEmergencyContactPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, err
	}

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	return &payload, nil
}
