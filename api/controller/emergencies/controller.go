package emergencies

import (
	"flavioltonon/hmv/application/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	usecases *Usecases
}

type Usecases struct {
	Emergencies usecases.EmergencyUsecase
}

func NewController(usecases *Usecases) *Controller {
	return &Controller{usecases: usecases}
}

func (c *Controller) SetRoutes(parent *mux.Router) {
	parent.HandleFunc("", c.createEmergency).Methods(http.MethodPost)
	parent.HandleFunc("", c.listEmergencies).Methods(http.MethodGet)
}
