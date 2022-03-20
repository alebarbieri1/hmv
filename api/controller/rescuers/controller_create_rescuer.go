package rescuers

import (
	"encoding/json"
	"net/http"

	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

func (c *Controller) createRescuer(w http.ResponseWriter, r *http.Request) {
	payload, err := c.newCreateRescuerRequestPayload(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToValidateRequest, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToValidateRequest, err))
		return
	}

	user, err := c.usecases.Users.FindUserByID(payload.UserID)
	if err == entity.ErrNotFound {
		c.drivers.Logger.Info(application.FailedToFindUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToFindUser, err))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToFindUser, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToFindUser, err))
		return
	}

	rescuer, err := c.usecases.Rescuers.CreateRescuer(user)
	if err == application.ErrUserAlreadyIsARescuer {
		c.drivers.Logger.Info(application.FailedToCreateRescuer, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToCreateRescuer, err))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToCreateRescuer, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToCreateRescuer, err))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewRescuer(rescuer)))
}

type createRescuerRequestPayload struct {
	UserID string `json:"user_id"`
}

func (p *createRescuerRequestPayload) Validate() error {
	return ozzo.ValidateStruct(p,
		ozzo.Field(&p.UserID, ozzo.Required),
	)
}

func (c *Controller) newCreateRescuerRequestPayload(r *http.Request) (*createRescuerRequestPayload, error) {
	var payload createRescuerRequestPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, err
	}

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	return &payload, nil
}
