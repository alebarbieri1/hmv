package api

import (
	"flavioltonon/hmv/api/controller"
	"flavioltonon/hmv/infrastructure/drivers"
	"flavioltonon/hmv/infrastructure/settings"
	"net/http"
)

// Server is the API server
type Server struct {
	core     *http.Server
	drivers  *drivers.Drivers
	settings *settings.Settings
}

func NewServer() (*Server, error) {
	settings, err := settings.New()
	if err != nil {
		return nil, err
	}

	drivers, err := drivers.New(settings)
	if err != nil {
		return nil, err
	}

	controller, err := controller.New(drivers)
	if err != nil {
		return nil, err
	}

	return &Server{
		core: &http.Server{
			Addr:    settings.Server.Address,
			Handler: controller.NewRouter(),
		},
		drivers:  drivers,
		settings: settings,
	}, nil
}

func (s *Server) Start() error {
	return s.core.ListenAndServe()
}

func (s *Server) Stop() error {
	return nil
}
