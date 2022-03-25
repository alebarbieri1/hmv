package repository

import (
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/infrastructure/repository/memory"
)

// Repositories defines all the repositories available in the application
type Repositories struct {
	Analysts    repositories.AnalystsRepository
	Emergencies repositories.EmergenciesRepository
	Pacients    repositories.PacientsRepository
	Rescuers    repositories.RescuersRepository
	Users       repositories.UsersRepository
}

// NewRepositories creates a new Repositories
func NewRepositories() (*Repositories, error) {
	analysts, err := memory.NewAnalystsRepository()
	if err != nil {
		return nil, err
	}

	emergencies, err := memory.NewEmergenciesRepository()
	if err != nil {
		return nil, err
	}

	pacients, err := memory.NewPacientsRepository()
	if err != nil {
		return nil, err
	}

	rescuers, err := memory.NewRescuersRepository()
	if err != nil {
		return nil, err
	}

	users, err := memory.NewUsersRepository()
	if err != nil {
		return nil, err
	}

	return &Repositories{
		Analysts:    analysts,
		Emergencies: emergencies,
		Pacients:    pacients,
		Rescuers:    rescuers,
		Users:       users,
	}, nil
}
