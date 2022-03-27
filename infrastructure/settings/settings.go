package settings

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Settings groups the application settings
type Settings struct {
	Server *ServerSettings

	Logging *LoggingSettings
}

type ServerSettings struct {
	// Address is the address of the server
	Address string

	// DevelopmentEnvironment defines if the server is running in a development environment
	DevelopmentEnvironment bool `mapstructure:"development_environment"`
}

type LoggingSettings struct {
	DevelopmentEnvironment bool `mapstructure:"development_environment"`
}

func init() {
	// Server settings
	viper.SetDefault("server.address", ":8080")
	viper.SetDefault("server.development_environment", true)

	// Logging settings
	viper.SetDefault("logging.development_environment", false)
}

// FromFile creates a new settings from a given file
func FromFile(path string) (*Settings, error) {
	var settings Settings

	// If the path to the file is not absolute, we should look for the file from the current working directory
	if !filepath.IsAbs(path) {
		workdir, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		path = filepath.Join(workdir, filepath.Clean(path))
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	// Set viper config file extension. Since filepath.Ext returns an extension preceded by a dot (e.g. ".yaml"),
	// we have to trim it manually.
	viper.SetConfigType(strings.TrimLeft(filepath.Ext(path), "."))

	// Read environment variables to get default values
	viper.AutomaticEnv()

	if err := viper.ReadConfig(f); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&settings); err != nil {
		return nil, err
	}

	return &settings, nil
}
