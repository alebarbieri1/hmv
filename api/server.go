package api

import (
	"flavioltonon/hmv/api/controller"
	"flavioltonon/hmv/infrastructure/drivers"
	"flavioltonon/hmv/infrastructure/repository"
	"flavioltonon/hmv/infrastructure/settings"
	"fmt"
	"net/http"
)

// Server is the API server
type Server struct {
	core         *http.Server
	repositories *repository.Repositories
	drivers      *drivers.Drivers
	settings     *settings.Settings
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

	return &Server{
		core: &http.Server{
			Addr:    settings.Server.Address,
			Handler: controller.New(drivers).NewRouter(),
		},
		drivers:  drivers,
		settings: settings,
	}, nil
}

func (s *Server) Start() error {
	s.drivers.Logger.Info(fmt.Sprintf("server listening and serving at %s", s.settings.Server.Address))
	return s.core.ListenAndServe()
}

func (s *Server) Stop() error {
	s.drivers.Logger.Info("server shutting down")
	return nil
}
