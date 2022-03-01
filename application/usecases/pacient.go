package usecases

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
)

type PacientUsecase interface {
	CreatePacient(user *entity.User) (*entity.Pacient, error)
	FindPacientByID(pacientID string) (*entity.Pacient, error)
	FindPacientByUserID(userID string) (*entity.Pacient, error)
	UpdateEmergencyContact(pacient *entity.Pacient, emergencyContact valueobject.EmergencyContact) (*entity.Pacient, error)
}
