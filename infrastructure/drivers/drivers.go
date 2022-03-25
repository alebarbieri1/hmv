package drivers

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/presenter"
	"flavioltonon/hmv/infrastructure/repository"
	"flavioltonon/hmv/infrastructure/settings"
	"time"
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

	if settings.Server.DevelopmentMode {
		repositories.Users.CreateUser(&entity.User{
			ID:       "af3ecfbc-0e18-4448-875a-d64744a1f5cd",
			Username: "paciente@teste.com",
			Password: "1234abc@",
			Data: valueobject.UserData{
				Name: "Lúcia Maria",
			},
			ProfileKind: valueobject.Pacient_ProfileKind,
			CreatedAt:   time.Date(2022, 03, 25, 21, 18, 33, 0, time.UTC),
			UpdatedAt:   time.Date(2022, 03, 25, 21, 18, 33, 0, time.UTC),
		})

		repositories.Pacients.CreatePacient(&entity.Pacient{
			ID:     "aed15ceb-c95e-4ee7-ae75-dd80e7f7da67",
			UserID: "af3ecfbc-0e18-4448-875a-d64744a1f5cd",
			EmergencyContact: valueobject.EmergencyContact{
				Name:         "Contato de emergência",
				MobileNumber: "5511999999999",
			},
			CreatedAt: time.Date(2022, 03, 25, 21, 18, 33, 0, time.UTC),
			UpdatedAt: time.Date(2022, 03, 25, 21, 18, 33, 0, time.UTC),
		})

		repositories.Users.CreateUser(&entity.User{
			ID:       "6d763219-d191-4970-85cf-5abeab25f9c4",
			Username: "analista@hmv.com.br",
			Password: "1234abc@",
			Data: valueobject.UserData{
				Name: "Ana Souza",
			},
			ProfileKind: valueobject.Analyst_ProfileKind,
			CreatedAt:   time.Date(2022, 03, 25, 21, 18, 33, 0, time.UTC),
			UpdatedAt:   time.Date(2022, 03, 25, 21, 18, 33, 0, time.UTC),
		})

		repositories.Analysts.CreateAnalyst(&entity.Analyst{
			ID:        "7470aade-d9b1-413e-8e76-d2eb837864ce",
			UserID:    "6d763219-d191-4970-85cf-5abeab25f9c4",
			CreatedAt: time.Date(2022, 03, 25, 21, 18, 33, 0, time.UTC),
			UpdatedAt: time.Date(2022, 03, 25, 21, 18, 33, 0, time.UTC),
		})

		repositories.Users.CreateUser(&entity.User{
			ID:       "16674382-3430-414e-b429-0e847999ca8e",
			Username: "socorrista@hmv.com.br",
			Password: "1234abc@",
			Data: valueobject.UserData{
				Name: "João Silva",
			},
			ProfileKind: valueobject.Rescuer_ProfileKind,
			CreatedAt:   time.Date(2022, 03, 25, 21, 18, 33, 0, time.UTC),
			UpdatedAt:   time.Date(2022, 03, 25, 21, 18, 33, 0, time.UTC),
		})

		repositories.Rescuers.CreateRescuer(&entity.Rescuer{
			ID:        "b6d41f8e-bf3c-451c-bc83-6f01312aed55",
			UserID:    "16674382-3430-414e-b429-0e847999ca8e",
			CreatedAt: time.Date(2022, 03, 25, 21, 18, 33, 0, time.UTC),
			UpdatedAt: time.Date(2022, 03, 25, 21, 18, 33, 0, time.UTC),
		})
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
