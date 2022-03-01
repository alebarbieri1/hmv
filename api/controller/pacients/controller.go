package pacients

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

	parent.HandleFunc("", c.createPacient).Methods(http.MethodPost)
	parent.HandleFunc("/emergency-contacts", c.updateEmergencyContact).Methods(http.MethodPut)
	parent.HandleFunc("/emergencies", c.listEmergencies).Methods(http.MethodGet)
	parent.HandleFunc("/emergencies", c.createEmergency).Methods(http.MethodPost)
}
