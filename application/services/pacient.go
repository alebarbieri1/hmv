package services

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
)

type PacientService struct {
	pacients repositories.PacientsRepository
}

func NewPacientService(repository repositories.PacientsRepository) (*PacientService, error) {
	return &PacientService{pacients: repository}, nil
}

func (s *PacientService) CreatePacient(userID string) (*entity.Pacient, error) {
	_, err := s.pacients.FindPacientByUserID(userID)
	if err == entity.ErrNotFound {
		pacient, err := entity.NewPacient(userID)
		if err != nil {
			return nil, err
		}

		if err := s.pacients.CreatePacient(pacient); err != nil {
			return nil, err
		}

		return pacient, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, entity.ErrUserAlreadyIsAPacient
}

func (s *PacientService) FindPacientByUserID(userID string) (*entity.Pacient, error) {
	return s.pacients.FindPacientByUserID(userID)
}
