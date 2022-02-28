package pacients

import (
	"net/http"

	"flavioltonon/hmv/application/usecases"
	"flavioltonon/hmv/infrastructure/drivers"

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
	parent.HandleFunc("", c.createPacient).Methods(http.MethodPost)
	parent.HandleFunc("/emergency-contacts", c.updateEmergencyContact).Methods(http.MethodPut)
	parent.HandleFunc("/emergencies", c.listEmergencies).Methods(http.MethodGet)
	parent.HandleFunc("/emergencies", c.createEmergency).Methods(http.MethodPost)
}
