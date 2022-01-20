package drivers

import (
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/infrastructure/repository/memory"
	"flavioltonon/hmv/infrastructure/settings"
)

// Drivers groups the application dependencies
type Drivers struct {
	Repositories *Repositories
}

// New creates new Drivers using a given set of Settings
func New(settings *settings.Settings) (*Drivers, error) {
	emergencies, err := memory.NewEmergenciesRepository()
	if err != nil {
		return nil, err
	}

	return &Drivers{
		Repositories: &Repositories{
			Emergencies: emergencies,
		},
	}, nil
}

type Repositories struct {
	Emergencies repositories.EmergenciesRepository
}
