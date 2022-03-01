package users

import (
	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) findUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := c.usecases.Users.FindUserByID(vars["user_id"])
	if err == entity.ErrNotFound {
		c.drivers.Logger.Info(application.FailedToFindUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.NotFound(application.FailedToFindUser, err))
		return
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToFindUser, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToFindUser, err))
		return
	}

	c.drivers.Presenter.Present(w, response.OK(presenter.NewUser(user)))
}
