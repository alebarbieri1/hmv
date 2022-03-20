package emergencies

import (
	"net/http"

	"flavioltonon/hmv/application/usecases"
	"flavioltonon/hmv/infrastructure/drivers"
	"flavioltonon/hmv/infrastructure/middleware"

	"github.com/gorilla/mux"
)

type Controller struct {
	usecases *Usecases
	drivers  *drivers.Drivers
}

type Usecases struct {
	Authentication usecases.AuthenticationUsecase
	Emergencies    usecases.EmergencyUsecase
	Pacients       usecases.PacientUsecase
}

func NewController(usecases *Usecases, drivers *drivers.Drivers) *Controller {
	return &Controller{usecases: usecases, drivers: drivers}
}

func (c *Controller) SetRoutes(parent *mux.Router) {
	parent.Use(mux.MiddlewareFunc(middleware.Authentication(
		c.usecases.Authentication,
		c.drivers.Logger,
		c.drivers.Presenter,
	)))

	parent.HandleFunc("", c.createEmergency).Methods(http.MethodPost)
	parent.HandleFunc("", c.listEmergencies).Methods(http.MethodGet)
	parent.HandleFunc("/{emergency_id}/form", c.updateEmergencyForm).Methods(http.MethodPut)
	parent.HandleFunc("/{emergency_id}/start", c.startEmergencyCare).Methods(http.MethodPatch)
	parent.HandleFunc("/{emergency_id}/cancel", c.cancelEmergency).Methods(http.MethodPatch)
}
