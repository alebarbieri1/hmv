package drivers

import (
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/presenter"
	"flavioltonon/hmv/infrastructure/settings"
)

// Drivers groups the application dependencies
type Drivers struct {
	Logger    logging.Logger
	Presenter presenter.Presenter
}

// New creates new Drivers using a given set of Settings
func New(settings *settings.Settings) (*Drivers, error) {
	logger, err := logging.NewZapLogger(settings.Logging)
	if err != nil {
		return nil, err
	}

	return &Drivers{
		Logger:    logger,
		Presenter: presenter.NewJSONPresenter(),
	}, nil
}
