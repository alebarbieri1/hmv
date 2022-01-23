package emergencies

import (
	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"
)

func (c *Controller) createEmergency(w http.ResponseWriter, r *http.Request) {
	user, err := c.usecases.Authentication.AuthenticateUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.ErrMsgFailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.ErrMsgFailedToAuthenticateUser, err))
		return
	}

	emergency, err := c.usecases.Emergencies.CreateEmergency(user.ID)
	if err != nil {
		c.drivers.Logger.Error(application.ErrMsgFailedToCreateEmergency, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.ErrMsgFailedToCreateEmergency, err))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewEmergency(emergency)))
}
