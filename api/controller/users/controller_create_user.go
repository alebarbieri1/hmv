package users

import (
	"encoding/json"
	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"
)

func (c *Controller) createUser(w http.ResponseWriter, r *http.Request) {
	payload, err := c.newCreateUserRequestPayload(r)
	if err != nil {
		c.drivers.Presenter.Present(w, response.BadRequest(err.Error()))
		return
	}

	user, err := c.usecases.Users.CreateUser(payload.Username, payload.Password)
	if err == entity.ErrUsernameAlreadyInUse {
		c.drivers.Presenter.Present(w, response.BadRequest(err.Error()))
		return
	}

	if err != nil {
		c.drivers.Presenter.Present(w, response.InternalServerError(err.Error()))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewUser(user)))
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
