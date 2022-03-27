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
	return &Repositories{
		Analysts:    memory.NewAnalystsRepository(),
		Emergencies: memory.NewEmergenciesRepository(),
		Pacients:    memory.NewPacientsRepository(),
		Rescuers:    memory.NewRescuersRepository(),
		Users:       memory.NewUsersRepository(),
	}, nil
}
