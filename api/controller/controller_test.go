package controller

import (
	"flavioltonon/hmv/api/controller/analysts"
	"flavioltonon/hmv/api/controller/emergencies"
	"flavioltonon/hmv/api/controller/pacients"
	"flavioltonon/hmv/api/controller/rescuers"
	"flavioltonon/hmv/api/controller/users"
	"flavioltonon/hmv/application/services"
	"flavioltonon/hmv/infrastructure/drivers"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/presenter"
	"flavioltonon/hmv/infrastructure/repository"
	"flavioltonon/hmv/infrastructure/repository/memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		drivers *drivers.Drivers
	}

	var (
		repositories = &repository.Repositories{
			Analysts:    memory.NewAnalystsRepository(),
			Emergencies: memory.NewEmergenciesRepository(),
			Pacients:    memory.NewPacientsRepository(),
			Rescuers:    memory.NewRescuersRepository(),
			Users:       memory.NewUsersRepository(),
		}

		logger = logging.NewNopLogger()

		presenter = presenter.NewJSONPresenter()
	)

	authenticationService := services.NewAuthenticationService(repositories.Users, logger)

	tests := []struct {
		name    string
		args    args
		want    *Controller
		wantErr bool
	}{
		{
			name: "Given a set of drivers.Drivers, a new Controller should be created",
			args: args{
				drivers: &drivers.Drivers{
					Repositories: repositories,
					Logger:       logger,
					Presenter:    presenter,
				},
			},
			want: &Controller{
				analysts: analysts.NewController(
					&analysts.Usecases{
						Authentication: authenticationService,
						Analysts: services.NewAnalystService(
							repositories.Analysts,
							repositories.Users,
							logger,
						),
					},
					&drivers.Drivers{
						Repositories: repositories,
						Logger:       logger,
						Presenter:    presenter,
					},
				),
				emergencies: emergencies.NewController(
					&emergencies.Usecases{
						Authentication: authenticationService,
						Emergencies: services.NewEmergencyService(
							repositories.Emergencies,
							repositories.Pacients,
							repositories.Users,
							logger,
						),
					},
					&drivers.Drivers{
						Repositories: repositories,
						Logger:       logger,
						Presenter:    presenter,
					},
				),
				pacients: pacients.NewController(
					&pacients.Usecases{
						Authentication: authenticationService,
						Pacients: services.NewPacientService(
							repositories.Pacients,
							repositories.Users,
							logger,
						),
					},
					&drivers.Drivers{
						Repositories: repositories,
						Logger:       logger,
						Presenter:    presenter,
					},
				),
				rescuers: rescuers.NewController(
					&rescuers.Usecases{
						Authentication: authenticationService,
						Rescuers: services.NewRescuerService(
							repositories.Rescuers,
							repositories.Users,
							logger,
						),
					},
					&drivers.Drivers{
						Repositories: repositories,
						Logger:       logger,
						Presenter:    presenter,
					},
				),
				users: users.NewController(
					&users.Usecases{
						Users: services.NewUserService(repositories.Users, logger),
					},
					&drivers.Drivers{
						Repositories: repositories,
						Logger:       logger,
						Presenter:    presenter,
					},
				),
				drivers: &drivers.Drivers{
					Repositories: repositories,
					Logger:       logger,
					Presenter:    presenter,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, New(tt.args.drivers))
		})
	}
}
