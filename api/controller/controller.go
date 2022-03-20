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
func New(drivers *drivers.Drivers) (*Controller, error) {
	authenticationService, err := services.NewAuthenticationService(drivers.Repositories.Users, drivers.Logger)
	if err != nil {
		return nil, err
	}

	analystsService, err := services.NewAnalystService(
		drivers.Repositories.Analysts,
		drivers.Repositories.Users,
		drivers.Logger,
	)
	if err != nil {
		return nil, err
	}

	emergenciesService, err := services.NewEmergencyService(drivers.Repositories.Emergencies, drivers.Repositories.Pacients, drivers.Logger)
	if err != nil {
		return nil, err
	}

	pacientsService, err := services.NewPacientService(
		drivers.Repositories.Pacients,
		drivers.Repositories.Users,
		drivers.Logger,
	)
	if err != nil {
		return nil, err
	}

	rescuersService, err := services.NewRescuerService(
		drivers.Repositories.Rescuers,
		drivers.Repositories.Users,
		drivers.Logger,
	)
	if err != nil {
		return nil, err
	}

	usersService, err := services.NewUserService(drivers.Repositories.Users, drivers.Logger)
	if err != nil {
		return nil, err
	}

	return &Controller{
		analysts: analysts.NewController(
			&analysts.Usecases{
				Authentication: authenticationService,
				Analysts:       analystsService,
				Users:          usersService,
			},
			drivers,
		),
		emergencies: emergencies.NewController(
			&emergencies.Usecases{
				Authentication: authenticationService,
				Emergencies:    emergenciesService,
				Pacients:       pacientsService,
			},
			drivers,
		),
		pacients: pacients.NewController(
			&pacients.Usecases{
				Authentication: authenticationService,
				Emergencies:    emergenciesService,
				Pacients:       pacientsService,
			},
			drivers,
		),
		rescuers: rescuers.NewController(
			&rescuers.Usecases{
				Authentication: authenticationService,
				Rescuers:       rescuersService,
				Users:          usersService,
			},
			drivers,
		),
		users: users.NewController(
			&users.Usecases{
				Users: usersService,
			},
			drivers,
		),
		drivers: drivers,
	}, nil
}

func (c *Controller) NewRouter() http.Handler {
	router := mux.NewRouter()
	c.analysts.SetRoutes(router.PathPrefix("/analysts").Subrouter())
	c.emergencies.SetRoutes(router.PathPrefix("/emergencies").Subrouter())
	c.pacients.SetRoutes(router.PathPrefix("/pacients").Subrouter())
	c.rescuers.SetRoutes(router.PathPrefix("/rescuers").Subrouter())
	c.users.SetRoutes(router.PathPrefix("/users").Subrouter())
	return alice.New(
		middleware.ResponseWrapper(),
		middleware.RequestID(),
		middleware.Logging(c.drivers.Logger),
	).Then(router)
}
