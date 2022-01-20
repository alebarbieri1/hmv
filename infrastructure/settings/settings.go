package settings

// Settings groups the application settings
type Settings struct {
	Server *ServerSettings
}

// New creates new Settings
func New() (*Settings, error) {
	return &Settings{
		Server: &ServerSettings{
			Address: ":8080",
		},
	}, nil
}

// ServerSettings are server settings
type ServerSettings struct {
	// Address is the address of the server
	Address string
}
