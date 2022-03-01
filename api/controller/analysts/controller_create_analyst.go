package analysts

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

func (c *Controller) createAnalyst(w http.ResponseWriter, r *http.Request) {
	payload, err := c.newCreateAnalystRequestPayload(r)
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

	analyst, err := c.usecases.Analysts.CreateAnalyst(user)
	if err == application.ErrUserAlreadyIsAnAnalyst {
		c.drivers.Logger.Info(application.FailedToCreateAnalyst, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToCreateAnalyst, err))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToCreateAnalyst, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToCreateAnalyst, err))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewAnalyst(analyst)))
}

type createAnalystRequestPayload struct {
	UserID string `json:"user_id"`
}

func (p *createAnalystRequestPayload) Validate() error {
	return ozzo.ValidateStruct(p,
		ozzo.Field(&p.UserID, ozzo.Required),
	)
}

func (c *Controller) newCreateAnalystRequestPayload(r *http.Request) (*createAnalystRequestPayload, error) {
	var payload createAnalystRequestPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, err
	}

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	return &payload, nil
}
