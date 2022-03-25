package pacients

import (
	"net/http"

	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
)

func (c *Controller) createPacient(w http.ResponseWriter, r *http.Request) {
	user, err := entity.NewUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.FailedToAuthenticateUser, err))
		return
	}

	pacient, err := c.usecases.Pacients.CreatePacient(user.ID)
	if err == application.ErrUserAlreadyIsAPacient {
		c.drivers.Logger.Debug(application.FailedToCreatePacient, logging.Error(err))
		c.drivers.Presenter.Present(w, response.OK(pacient))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToCreatePacient, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToCreatePacient, err))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewPacient(pacient)))
}
