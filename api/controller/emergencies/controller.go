package emergencies

import (
	"flavioltonon/hmv/application/usecases"
	"flavioltonon/hmv/infrastructure/drivers"
	"net/http"

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
	parent.HandleFunc("", c.createEmergency).Methods(http.MethodPost)
	parent.HandleFunc("", c.listEmergencies).Methods(http.MethodGet)
}
