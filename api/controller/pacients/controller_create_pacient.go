package pacients

import (
	"encoding/json"
	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"
)

func (c *Controller) createPacient(w http.ResponseWriter, r *http.Request) {
	username, password, hasCredentials := r.BasicAuth()
	if !hasCredentials {
		c.drivers.Presenter.Present(w, response.Unauthorized("basic authentication is required"))
		return
	}

	user, err := c.usecases.Authentication.AuthenticateUser(username, password)
	if err != nil {
		c.drivers.Presenter.Present(w, response.Unauthorized(err.Error()))
		return
	}

	pacient, err := c.usecases.Pacients.CreatePacient(user.ID)
	if err == entity.ErrUserAlreadyIsAPacient {
		c.drivers.Presenter.Present(w, response.BadRequest(err.Error()))
		return
	}

	if err != nil {
		c.drivers.Presenter.Present(w, response.InternalServerError(err.Error()))
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
