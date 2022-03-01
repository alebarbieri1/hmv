package pacients

import (
	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) findPacient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pacient, err := c.usecases.Pacients.FindPacientByID(vars["pacient_id"])
	if err == entity.ErrNotFound {
		c.drivers.Logger.Info(application.FailedToFindPacient, logging.Error(err))
		c.drivers.Presenter.Present(w, response.NotFound(application.FailedToFindPacient, err))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToFindPacient, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToFindPacient, err))
		return
	}

	c.drivers.Presenter.Present(w, response.OK(presenter.NewPacient(pacient)))
}
