package drivers

import (
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/presenter"
	"flavioltonon/hmv/infrastructure/repository"
	"flavioltonon/hmv/infrastructure/settings"
)

// Drivers groups the application dependencies
type Drivers struct {
	Repositories *repository.Repositories
	Logger       logging.Logger
	Presenter    presenter.Presenter
}

// New creates new Drivers using a given set of Settings
func New(settings *settings.Settings) (*Drivers, error) {
	repositories, err := repository.NewRepositories()
	if err != nil {
		return nil, err
	}

	logger, err := logging.NewZapLogger(settings.Logging)
	if err != nil {
		return nil, err
	}

	return &Drivers{
		Repositories: repositories,
		Logger:       logger,
		Presenter:    presenter.NewJSONPresenter(),
	}, nil
}

func (d *Drivers) Stop() error {
	return nil
}
