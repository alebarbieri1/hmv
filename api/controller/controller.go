package controller

import (
	"flavioltonon/hmv/api/controller/emergencies"
	"flavioltonon/hmv/application/services"
	"flavioltonon/hmv/infrastructure/drivers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Controller is the application controller
type Controller struct {
	emergencies *emergencies.Controller
}

// New creates a new Controller with a given set of Drivers
func New(drivers *drivers.Drivers) (*Controller, error) {
	emergenciesService, err := services.NewEmergencyService(drivers.Repositories.Emergencies)
	if err != nil {
		return nil, err
	}

	return &Controller{
		emergencies: emergencies.NewController(&emergencies.Usecases{
			Emergencies: emergenciesService,
		}),
	}, nil
}

func (c *Controller) NewRouter() http.Handler {
	router := mux.NewRouter()
	c.emergencies.SetRoutes(router.PathPrefix("/emergencies").Subrouter())

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		template, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		fmt.Printf("template: %v\n", template)

		return nil
	})

	return router
}
