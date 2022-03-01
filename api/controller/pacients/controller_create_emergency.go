package pacients

import (
	"net/http"

	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
)

func (c *Controller) createEmergency(w http.ResponseWriter, r *http.Request) {
	user, err := entity.NewUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.FailedToAuthenticateUser, err))
		return
	}

	if !user.IsPacient() {
		c.drivers.Logger.Info(application.FailedToCreateEmergency, logging.Error(application.ErrUserMustBeAPacient))
		c.drivers.Presenter.Present(w, response.Forbidden(application.FailedToCreateEmergency, application.ErrUserMustBeAPacient))
		return
	}

	emergency, err := c.usecases.Emergencies.CreateEmergency(user.ID)
	if err != nil {
		c.drivers.Logger.Error(application.FailedToCreateEmergency, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToCreateEmergency, err))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewEmergency(emergency)))
}
