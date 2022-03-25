package rescuers

import (
	"net/http"

	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
)

func (c *Controller) createRescuer(w http.ResponseWriter, r *http.Request) {
	user, err := entity.NewUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.FailedToAuthenticateUser, err))
		return
	}

	rescuer, err := c.usecases.Rescuers.CreateRescuer(user.ID)
	if err == application.ErrUserAlreadyIsARescuer {
		c.drivers.Logger.Info(application.FailedToCreateRescuer, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToCreateRescuer, err))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToCreateRescuer, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToCreateRescuer, err))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewRescuer(rescuer)))
}
