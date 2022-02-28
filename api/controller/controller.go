package controller

import (
	"net/http"

	"flavioltonon/hmv/api/controller/pacients"
	"flavioltonon/hmv/api/controller/users"
	"flavioltonon/hmv/application/services"
	"flavioltonon/hmv/infrastructure/drivers"
	"flavioltonon/hmv/infrastructure/middleware"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Controller is the application controller
type Controller struct {
	pacients *pacients.Controller
	users    *users.Controller
	drivers  *drivers.Drivers
}

// New creates a new Controller with a given set of Drivers
func New(drivers *drivers.Drivers) (*Controller, error) {
	authenticationService, err := services.NewAuthenticationService(drivers.Repositories.Users, drivers.Logger)
	if err != nil {
		return nil, err
	}

	emergenciesService, err := services.NewEmergencyService(drivers.Repositories.Emergencies, drivers.Repositories.Pacients, drivers.Logger)
	if err != nil {
		return nil, err
	}

	pacientsService, err := services.NewPacientService(drivers.Repositories.Pacients, drivers.Logger)
	if err != nil {
		return nil, err
	}

	usersService, err := services.NewUserService(drivers.Repositories.Users, drivers.Logger)
	if err != nil {
		return nil, err
	}

	return &Controller{
		pacients: pacients.NewController(
			&pacients.Usecases{
				Authentication: authenticationService,
				Emergencies:    emergenciesService,
				Pacients:       pacientsService,
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
	c.pacients.SetRoutes(router.PathPrefix("/pacients").Subrouter())
	c.users.SetRoutes(router.PathPrefix("/users").Subrouter())
	return alice.New(
		middleware.ResponseWrapper(),
		middleware.RequestID(),
		middleware.Logging(c.drivers.Logger),
	).Then(router)
}
