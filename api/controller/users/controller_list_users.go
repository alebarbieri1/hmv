package users

import (
	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"
)

func (c *Controller) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.usecases.Users.ListUsers()
	if err != nil {
		c.drivers.Logger.Error(application.FailedToListUsers, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToListUsers, err))
		return
	}

	c.drivers.Presenter.Present(w, response.OK(presenter.NewUsers(users)))
}
