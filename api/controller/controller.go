package controller

import (
	"flavioltonon/hmv/api/controller/emergencies"
	"flavioltonon/hmv/api/controller/pacients"
	"flavioltonon/hmv/api/controller/users"
	"flavioltonon/hmv/application/services"
	"flavioltonon/hmv/infrastructure/drivers"
	"net/http"

	"github.com/gorilla/mux"
)

// Controller is the application controller
type Controller struct {
	emergencies *emergencies.Controller
	pacients    *pacients.Controller
	users       *users.Controller
}

// New creates a new Controller with a given set of Drivers
func New(drivers *drivers.Drivers) (*Controller, error) {
	authenticationService, err := services.NewAuthenticationService(drivers.Repositories.Users)
	if err != nil {
		return nil, err
	}

	emergenciesService, err := services.NewEmergencyService(drivers.Repositories.Emergencies)
	if err != nil {
		return nil, err
	}

	pacientsService, err := services.NewPacientService(drivers.Repositories.Pacients)
	if err != nil {
		return nil, err
	}

	usersService, err := services.NewUserService(drivers.Repositories.Users)
	if err != nil {
		return nil, err
	}

	return &Controller{
		emergencies: emergencies.NewController(
			&emergencies.Usecases{
				Authentication: authenticationService,
				Emergencies:    emergenciesService,
				Pacients:       pacientsService,
			},
			&emergencies.Drivers{
				Presenter: drivers.Presenter,
			},
		),
		pacients: pacients.NewController(
			&pacients.Usecases{
				Authentication: authenticationService,
				Pacients:       pacientsService,
			},
			&pacients.Drivers{
				Presenter: drivers.Presenter,
			},
		),
		users: users.NewController(
			&users.Usecases{
				Users: usersService,
			},
			&users.Drivers{
				Presenter: drivers.Presenter,
			},
		),
	}, nil
}

func (c *Controller) NewRouter() http.Handler {
	router := mux.NewRouter()
	c.emergencies.SetRoutes(router.PathPrefix("/emergencies").Subrouter())
	c.pacients.SetRoutes(router.PathPrefix("/pacients").Subrouter())
	c.users.SetRoutes(router.PathPrefix("/users").Subrouter())
	return router
}
