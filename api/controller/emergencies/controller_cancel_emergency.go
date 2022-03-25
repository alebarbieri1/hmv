package emergencies

import (
	"net/http"

	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/infrastructure/response"

	"github.com/gorilla/mux"
)

func (c *Controller) cancelEmergency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	emergency, err := c.usecases.Emergencies.CancelEmergency(vars["emergency_id"])
	if err != nil {
		c.drivers.Logger.Error(application.FailedToUpdateEmergency, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToUpdateEmergency, err))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewEmergency(emergency)))
}
