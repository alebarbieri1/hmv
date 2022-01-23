package pacients

import (
	"flavioltonon/hmv/application/usecases"
	"flavioltonon/hmv/infrastructure/presenter"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	usecases *Usecases
	drivers  *Drivers
}

type Usecases struct {
	Authentication usecases.AuthenticationUsecase
	Pacients       usecases.PacientUsecase
}

type Drivers struct {
	Presenter presenter.Presenter
}

func NewController(usecases *Usecases, drivers *Drivers) *Controller {
	return &Controller{usecases: usecases, drivers: drivers}
}

func (c *Controller) SetRoutes(parent *mux.Router) {
	parent.HandleFunc("", c.createPacient).Methods(http.MethodPost)
}
