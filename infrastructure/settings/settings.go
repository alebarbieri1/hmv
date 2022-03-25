package settings

import "flavioltonon/hmv/infrastructure/logging"

// Settings groups the application settings
type Settings struct {
	Server  *ServerSettings
	Logging *logging.Settings
}

// New creates new Settings
func New() (*Settings, error) {
	return &Settings{
		Server: &ServerSettings{
			Address:         ":8080",
			DevelopmentMode: true,
		},
		Logging: &logging.Settings{
			DevelopmentMode: true,
		},
	}, nil
}

// ServerSettings are server settings
type ServerSettings struct {
	// Address is the address of the server
	Address string

	// DevelopmentMode defines if the server is running in development mode
	DevelopmentMode bool
}
