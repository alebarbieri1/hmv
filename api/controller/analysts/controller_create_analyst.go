package analysts

import (
	"net/http"

	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
)

func (c *Controller) createAnalyst(w http.ResponseWriter, r *http.Request) {
	user, err := entity.NewUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.FailedToAuthenticateUser, err))
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
