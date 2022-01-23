package repositories

import "flavioltonon/hmv/domain/entity"

type PacientsRepository interface {
	CreatePacient(pacient *entity.Pacient) error
	FindPacientByUserID(userID string) (*entity.Pacient, error)
}
