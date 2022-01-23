package users

import (
	"encoding/json"
	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"
)

func (c *Controller) createUser(w http.ResponseWriter, r *http.Request) {
	payload, err := c.newCreateUserRequestPayload(r)
	if err != nil {
		c.drivers.Logger.Info(application.ErrMsgFailedToValidateRequest, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.ErrMsgFailedToValidateRequest, err))
		return
	}

	user, err := c.usecases.Users.CreateUser(payload.Username, payload.Password)
	if err == application.ErrUsernameAlreadyInUse {
		c.drivers.Logger.Info(application.ErrMsgFailedToCreateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.ErrMsgFailedToCreateUser, err))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.ErrMsgFailedToCreateUser, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.ErrMsgFailedToCreateUser, err))
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
