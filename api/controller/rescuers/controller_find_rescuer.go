package rescuers

import (
	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) findRescuer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	rescuer, err := c.usecases.Rescuers.FindRescuerByID(vars["rescuer_id"])
	if err == entity.ErrNotFound {
		c.drivers.Logger.Info(application.FailedToFindRescuer, logging.Error(err))
		c.drivers.Presenter.Present(w, response.NotFound(application.FailedToFindRescuer, err))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToFindRescuer, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToFindRescuer, err))
		return
	}

	c.drivers.Presenter.Present(w, response.OK(presenter.NewRescuer(rescuer)))
}
