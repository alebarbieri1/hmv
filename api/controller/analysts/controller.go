package analysts

import (
	"net/http"

	"flavioltonon/hmv/application/usecases"
	"flavioltonon/hmv/infrastructure/drivers"
	"flavioltonon/hmv/infrastructure/middleware"

	"github.com/gorilla/mux"
)

// Controller is an analysts service controller
type Controller struct {
	usecases *Usecases
	drivers  *drivers.Drivers
}

// Usecases are the application usecases used by the Controller
type Usecases struct {
	Authentication usecases.AuthenticationUsecase
	Analysts       usecases.AnalystUsecase
}

// NewController creates a new Controller
func NewController(usecases *Usecases, drivers *drivers.Drivers) *Controller {
	return &Controller{usecases: usecases, drivers: drivers}
}

// SetRoutes sets the Controller routes to a given mux.Router
func (c *Controller) SetRoutes(parent *mux.Router) {
	parent.Use(mux.MiddlewareFunc(middleware.Authentication(
		c.usecases.Authentication,
		c.drivers.Logger,
		c.drivers.Presenter,
	)))

	parent.HandleFunc("", c.createAnalyst).Methods(http.MethodPost)
	parent.HandleFunc("/{analyst_id}", c.findAnalyst).Methods(http.MethodGet)
}
