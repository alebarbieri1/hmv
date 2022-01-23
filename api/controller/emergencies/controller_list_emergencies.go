package emergencies

import (
	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"
)

func (c *Controller) listEmergencies(w http.ResponseWriter, r *http.Request) {
	user, err := c.usecases.Authentication.AuthenticateUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.ErrMsgFailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.ErrMsgFailedToAuthenticateUser, err))
		return
	}

	pacient, err := c.usecases.Pacients.FindPacientByUserID(user.ID)
	if err == entity.ErrNotFound {
		c.drivers.Logger.Info(application.ErrMsgFailedToFindPacient, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Forbidden(application.ErrMsgFailedToListEmergencies, application.ErrUserMustBeAPacient))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.ErrMsgFailedToFindPacient, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.ErrMsgFailedToListEmergencies, err))
		return
	}

	emergencies, err := c.usecases.Emergencies.ListEmergenciesByPacientID(pacient.ID)
	if err != nil {
		c.drivers.Logger.Error(application.ErrMsgFailedToListEmergencies, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.ErrMsgFailedToListEmergencies, err))
		return
	}

	c.drivers.Presenter.Present(w, response.OK(presenter.NewEmergencies(emergencies)))
}
