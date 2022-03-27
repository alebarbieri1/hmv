package controller

import (
	"net/http"

	"flavioltonon/hmv/api/controller/analysts"
	"flavioltonon/hmv/api/controller/emergencies"
	"flavioltonon/hmv/api/controller/pacients"
	"flavioltonon/hmv/api/controller/rescuers"
	"flavioltonon/hmv/api/controller/users"
	"flavioltonon/hmv/application/services"
	"flavioltonon/hmv/infrastructure/drivers"
	"flavioltonon/hmv/infrastructure/middleware"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Controller is the application controller
type Controller struct {
	analysts    *analysts.Controller
	emergencies *emergencies.Controller
	pacients    *pacients.Controller
	rescuers    *rescuers.Controller
	users       *users.Controller
	drivers     *drivers.Drivers
}

// New creates a new Controller with a given set of Drivers
func New(drivers *drivers.Drivers) *Controller {
	authenticationService := services.NewAuthenticationService(drivers.Repositories.Users, drivers.Logger)

	return &Controller{
		analysts: analysts.NewController(
			&analysts.Usecases{
				Authentication: authenticationService,
				Analysts: services.NewAnalystService(
					drivers.Repositories.Analysts,
					drivers.Repositories.Users,
					drivers.Logger,
				),
			},
			drivers,
		),
		emergencies: emergencies.NewController(
			&emergencies.Usecases{
				Authentication: authenticationService,
				Emergencies: services.NewEmergencyService(
					drivers.Repositories.Emergencies,
					drivers.Repositories.Pacients,
					drivers.Repositories.Users,
					drivers.Logger,
				),
			},
			drivers,
		),
		pacients: pacients.NewController(
			&pacients.Usecases{
				Authentication: authenticationService,
				Pacients: services.NewPacientService(
					drivers.Repositories.Pacients,
					drivers.Repositories.Users,
					drivers.Logger,
				),
			},
			drivers,
		),
		rescuers: rescuers.NewController(
			&rescuers.Usecases{
				Authentication: authenticationService,
				Rescuers: services.NewRescuerService(
					drivers.Repositories.Rescuers,
					drivers.Repositories.Users,
					drivers.Logger,
				),
			},
			drivers,
		),
		users: users.NewController(
			&users.Usecases{
				Users: services.NewUserService(drivers.Repositories.Users, drivers.Logger),
			},
			drivers,
		),
		drivers: drivers,
	}
}

// NewRouter creates a new http.Handler with the controller routes
func (c *Controller) NewRouter() http.Handler {
	router := mux.NewRouter()
	c.analysts.SetRoutes(router.PathPrefix("/analysts").Subrouter())
	c.emergencies.SetRoutes(router.PathPrefix("/emergencies").Subrouter())
	c.pacients.SetRoutes(router.PathPrefix("/pacients").Subrouter())
	c.rescuers.SetRoutes(router.PathPrefix("/rescuers").Subrouter())
	c.users.SetRoutes(router.PathPrefix("/users").Subrouter())
	return alice.New(
		middleware.CORS(),
		middleware.ResponseWrapper(),
		middleware.RequestID(),
		middleware.Logging(c.drivers.Logger),
	).Then(router)
}
