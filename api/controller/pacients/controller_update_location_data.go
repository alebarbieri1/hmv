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

func (c *Controller) updateLocationData(w http.ResponseWriter, r *http.Request) {
	user, err := entity.NewUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.FailedToAuthenticateUser, err))
		return
	}

	payload, err := c.newUpdateLocationDataPayload(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToValidateRequest, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToValidateRequest, err))
		return
	}

	vars := mux.Vars(r)

	updatedPacient, err := c.usecases.Pacients.UpdateLocationData(user.ID, vars["pacient_id"], payload.toValueObject())
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

type updateLocationDataPayload struct {
	State   string `json:"state"`
	City    string `json:"city"`
	Address string `json:"address"`
	ZipCode string `json:"zipcode"`
}

func (p *updateLocationDataPayload) Validate() error {
	return ozzo.ValidateStruct(p,
		ozzo.Field(&p.State, ozzo.Required),
		ozzo.Field(&p.City, ozzo.Required),
		ozzo.Field(&p.Address, ozzo.Required),
		ozzo.Field(&p.ZipCode, ozzo.Required),
	)
}

func (p *updateLocationDataPayload) toValueObject() valueobject.LocationData {
	return valueobject.LocationData{
		State:   p.State,
		City:    p.City,
		Address: p.Address,
		ZipCode: p.ZipCode,
	}
}

func (c *Controller) newUpdateLocationDataPayload(r *http.Request) (*updateLocationDataPayload, error) {
	var payload updateLocationDataPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, err
	}

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	return &payload, nil
}
