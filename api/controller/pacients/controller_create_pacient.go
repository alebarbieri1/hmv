package pacients

import (
	"encoding/json"
	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"
)

func (c *Controller) createPacient(w http.ResponseWriter, r *http.Request) {
	user, err := c.usecases.Authentication.AuthenticateUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.ErrMsgFailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.ErrMsgFailedToAuthenticateUser, err))
		return
	}

	pacient, err := c.usecases.Pacients.CreatePacient(user.ID)
	if err == application.ErrUserAlreadyIsAPacient {
		c.drivers.Logger.Info(application.ErrMsgFailedToCreatePacient, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.ErrMsgFailedToCreatePacient, err))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.ErrMsgFailedToCreatePacient, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.ErrMsgFailedToCreatePacient, err))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewPacient(pacient)))
}

type createUserRequestPayload struct {
	Username string
	Password string
}

func (c *Controller) newCreateUserRequestPayload(r *http.Request) (*createUserRequestPayload, error) {
	var payload createUserRequestPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
