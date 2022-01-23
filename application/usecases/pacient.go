package usecases

import "flavioltonon/hmv/domain/entity"

type PacientUsecase interface {
	CreatePacient(userID string) (*entity.Pacient, error)
	FindPacientByUserID(userID string) (*entity.Pacient, error)
}
