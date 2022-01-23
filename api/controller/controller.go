package controller

import (
	"net/http"

	"flavioltonon/hmv/api/controller/emergencies"
	"flavioltonon/hmv/api/controller/pacients"
	"flavioltonon/hmv/api/controller/users"
	"flavioltonon/hmv/application/services"
	"flavioltonon/hmv/infrastructure/drivers"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/middleware"
	"flavioltonon/hmv/infrastructure/repository"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Controller is the application controller
type Controller struct {
	emergencies *emergencies.Controller
	pacients    *pacients.Controller
	users       *users.Controller
	drivers     *drivers.Drivers
}

// New creates a new Controller with a given set of Drivers
func New(repositories *repository.Repositories, drivers *drivers.Drivers) (*Controller, error) {
	authenticationService, err := services.NewAuthenticationService(repositories.Users)
	if err != nil {
		return nil, err
	}

	emergenciesService, err := services.NewEmergencyService(repositories.Emergencies, repositories.Pacients)
	if err != nil {
		return nil, err
	}

	pacientsService, err := services.NewPacientService(repositories.Pacients)
	if err != nil {
		return nil, err
	}

	usersService, err := services.NewUserService(repositories.Users)
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
			drivers,
		),
		pacients: pacients.NewController(
			&pacients.Usecases{
				Authentication: authenticationService,
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
	c.emergencies.SetRoutes(router.PathPrefix("/emergencies").Subrouter())
	c.pacients.SetRoutes(router.PathPrefix("/pacients").Subrouter())
	c.users.SetRoutes(router.PathPrefix("/users").Subrouter())
	return alice.New(
		middleware.RequestID(),
		logging.Middleware(c.drivers.Logger),
	).Then(router)
}
