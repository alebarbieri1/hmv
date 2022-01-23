package drivers

import (
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/infrastructure/presenter"
	"flavioltonon/hmv/infrastructure/repository/memory"
	"flavioltonon/hmv/infrastructure/settings"
)

// Drivers groups the application dependencies
type Drivers struct {
	Presenter    presenter.Presenter
	Repositories *Repositories
}

// New creates new Drivers using a given set of Settings
func New(settings *settings.Settings) (*Drivers, error) {
	emergencies, err := memory.NewEmergenciesRepository()
	if err != nil {
		return nil, err
	}

	pacients, err := memory.NewPacientsRepository()
	if err != nil {
		return nil, err
	}

	users, err := memory.NewUsersRepository()
	if err != nil {
		return nil, err
	}

	return &Drivers{
		Presenter: presenter.NewJSONPresenter(),
		Repositories: &Repositories{
			Emergencies: emergencies,
			Pacients:    pacients,
			Users:       users,
		},
	}, nil
}

type Repositories struct {
	Emergencies repositories.EmergenciesRepository
	Pacients    repositories.PacientsRepository
	Users       repositories.UsersRepository
}
