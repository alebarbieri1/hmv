package repository

import (
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/infrastructure/repository/memory"
)

type Repositories struct {
	Emergencies repositories.EmergenciesRepository
	Pacients    repositories.PacientsRepository
	Users       repositories.UsersRepository
}

func NewRepositories() (*Repositories, error) {
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

	return &Repositories{
		Emergencies: emergencies,
		Pacients:    pacients,
		Users:       users,
	}, nil
}
