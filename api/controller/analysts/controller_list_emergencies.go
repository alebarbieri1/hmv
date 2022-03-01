package analysts

import (
	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"
)

func (c *Controller) listEmergencies(w http.ResponseWriter, r *http.Request) {
	user, err := entity.NewUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.FailedToAuthenticateUser, err))
		return
	}

	if !user.IsAnalyst() {
		c.drivers.Logger.Info(application.FailedToListEmergencies, logging.Error(application.ErrUserMustBeAnAnalyst))
		c.drivers.Presenter.Present(w, response.Forbidden(application.FailedToListEmergencies, application.ErrUserMustBeAnAnalyst))
		return
	}

	var emergencies []*entity.Emergency

	s := r.URL.Query().Get("status")

	if status := valueobject.EmergencyStatusFromString(s); status == valueobject.Undefined_EmergencyStatus {
		emergencies, err = c.usecases.Emergencies.ListEmergencies()
	} else {
		emergencies, err = c.usecases.Emergencies.ListEmergenciesByStatus(status)
	}

	if err != nil {
		c.drivers.Logger.Error(application.FailedToListEmergencies, err, logging.String("status", s))
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToListEmergencies, err))
		return
	}

	c.drivers.Logger.Debug("emergencies listed", logging.String("user_id", user.ID))
	c.drivers.Presenter.Present(w, response.OK(presenter.NewEmergencies(emergencies)))
}
