package pacients

import (
	"encoding/json"
	"net/http"

	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
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

	if !user.IsPacient() {
		c.drivers.Logger.Info(application.FailedToUpdatePacient, logging.Error(application.ErrUserMustBeAPacient))
		c.drivers.Presenter.Present(w, response.Forbidden(application.FailedToUpdatePacient, application.ErrUserMustBeAPacient))
		return
	}

	pacient, err := c.usecases.Pacients.FindPacientByUserID(user.ID)
	if err == entity.ErrNotFound {
		c.drivers.Logger.Info(application.FailedToFindPacient, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Forbidden(application.FailedToListEmergencies, application.ErrUserMustBeAPacient))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToFindPacient, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToListEmergencies, err))
		return
	}

	updatedPacient, err := c.usecases.Pacients.UpdateEmergencyContact(pacient.ID, payload.toValueObject())
	if err == application.ErrUserAlreadyIsAPacient {
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
