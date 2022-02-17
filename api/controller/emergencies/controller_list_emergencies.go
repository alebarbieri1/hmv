package emergencies

import (
	"net/http"

	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
)

func (c *Controller) listEmergencies(w http.ResponseWriter, r *http.Request) {
	user, err := c.usecases.Authentication.AuthenticateUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.FailedToAuthenticateUser, err))
		return
	}

	pacient, err := c.usecases.Pacients.FindPacientByUserID(user.ID)
	if err == entity.ErrNotFound {
		c.drivers.Logger.Info(application.FailedToFindPacient, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Forbidden(application.FailedToListEmergencies, application.ErrUserMustBeAPacient))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToFindPacient, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToListEmergencies, err))
		return
	}

	emergencies, err := c.usecases.Emergencies.ListEmergenciesByPacientID(pacient.ID)
	if err != nil {
		c.drivers.Logger.Error(application.FailedToListEmergencies, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToListEmergencies, err))
		return
	}

	c.drivers.Presenter.Present(w, response.OK(presenter.NewEmergencies(emergencies)))
}
