package usecases

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
)

type PacientUsecase interface {
	CreatePacient(userID string) (*entity.Pacient, error)
	FindPacientByUserID(userID string) (*entity.Pacient, error)
	UpdateEmergencyContact(userID string, emergencyContact valueobject.EmergencyContact) (*entity.Pacient, error)
}
