package emergencies

import (
	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"
)

func (c *Controller) listEmergencies(w http.ResponseWriter, r *http.Request) {
	username, password, hasCredentials := r.BasicAuth()
	if !hasCredentials {
		c.drivers.Presenter.Present(w, response.Unauthorized("basic authentication is required"))
		return
	}

	user, err := c.usecases.Authentication.AuthenticateUser(username, password)
	if err != nil {
		c.drivers.Presenter.Present(w, response.Unauthorized(err.Error()))
		return
	}

	pacient, err := c.usecases.Pacients.FindPacientByUserID(user.ID)
	if err == entity.ErrNotFound {
		c.drivers.Presenter.Present(w, response.Forbidden("user must be a pacient"))
		return
	}

	if err != nil {
		c.drivers.Presenter.Present(w, response.InternalServerError(err.Error()))
		return
	}

	emergencies, err := c.usecases.Emergencies.ListEmergenciesByPacientID(pacient.ID)
	if err != nil {
		c.drivers.Presenter.Present(w, response.InternalServerError(err.Error()))
		return
	}

	c.drivers.Presenter.Present(w, response.OK(presenter.NewEmergencies(emergencies)))
}
