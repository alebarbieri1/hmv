package emergencies

import (
	"net/http"

	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"

	"github.com/gorilla/mux"
)

func (c *Controller) removePacient(w http.ResponseWriter, r *http.Request) {
	user, err := entity.NewUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.FailedToAuthenticateUser, err))
		return
	}

	vars := mux.Vars(r)

	emergency, err := c.usecases.Emergencies.FindEmergencyByID(vars["emergency_id"])
	if err == entity.ErrNotFound {
		c.drivers.Logger.Info(application.FailedToFindEmergency, logging.Error(err))
		c.drivers.Presenter.Present(w, response.NotFound(application.FailedToFindEmergency, err))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToFindEmergency, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToFindEmergency, err))
		return
	}

	err = c.usecases.Emergencies.RemovePacient(user, emergency)
	if err == application.ErrInternalError {
		c.drivers.Logger.Error(application.FailedToUpdateEmergency, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToUpdateEmergency, err))
		return
	}
	if err == application.ErrUserMustBeARescuer {
		c.drivers.Logger.Error(application.FailedToUpdateEmergency, err)
		c.drivers.Presenter.Present(w, response.Forbidden(application.FailedToUpdateEmergency, err))
		return
	}
	if err != nil {
		c.drivers.Logger.Error(application.FailedToUpdateEmergency, err)
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToUpdateEmergency, err))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewEmergency(emergency)))
}
