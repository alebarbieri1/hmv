package users

import (
	"encoding/json"
	"net/http"

	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

func (c *Controller) createUser(w http.ResponseWriter, r *http.Request) {
	payload, err := c.newCreateUserRequestPayload(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToValidateRequest, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToValidateRequest, err))
		return
	}

	user, err := c.usecases.Users.CreateUser(payload.Username, payload.Password)
	if err == application.ErrUsernameAlreadyInUse {
		c.drivers.Logger.Info(application.FailedToCreateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToCreateUser, err))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToCreateUser, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToCreateUser, err))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewUser(user)))
}

type createUserRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p *createUserRequestPayload) Validate() error {
	return ozzo.ValidateStruct(p,
		ozzo.Field(&p.Username, ozzo.Required),
		ozzo.Field(&p.Password, ozzo.Required),
	)
}

func (c *Controller) newCreateUserRequestPayload(r *http.Request) (*createUserRequestPayload, error) {
	var payload createUserRequestPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, err
	}

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	return &payload, nil
}