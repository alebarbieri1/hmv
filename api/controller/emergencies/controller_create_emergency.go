package emergencies

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

	emergency, err := c.usecases.Emergencies.CreateEmergency(user.ID)
	if err == application.ErrInternalError {
		c.drivers.Logger.Error(application.FailedToCreateEmergency, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToCreateEmergency, err))
		return
	}
	if err == application.ErrUserMustBeAPacient {
		c.drivers.Logger.Error(application.FailedToCreateEmergency, err)
		c.drivers.Presenter.Present(w, response.Forbidden(application.FailedToCreateEmergency, err))
		return
	}
	if err != nil {
		c.drivers.Logger.Error(application.FailedToCreateEmergency, err)
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToCreateEmergency, err))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewEmergency(emergency)))
}
