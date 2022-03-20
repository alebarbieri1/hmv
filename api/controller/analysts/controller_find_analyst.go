package analysts

import (
	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) findAnalyst(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	analyst, err := c.usecases.Analysts.FindAnalystByID(vars["analyst_id"])
	if err == entity.ErrNotFound {
		c.drivers.Logger.Info(application.FailedToFindAnalyst, logging.Error(err))
		c.drivers.Presenter.Present(w, response.NotFound(application.FailedToFindAnalyst, err))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToFindAnalyst, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToFindAnalyst, err))
		return
	}

	c.drivers.Presenter.Present(w, response.OK(presenter.NewAnalyst(analyst)))
}
