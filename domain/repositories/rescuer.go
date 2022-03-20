package repositories

import "flavioltonon/hmv/domain/entity"

type RescuersRepository interface {
	CreateRescuer(pacient *entity.Rescuer) error
	FindRescuerByID(rescuerID string) (*entity.Rescuer, error)
	FindRescuerByUserID(userID string) (*entity.Rescuer, error)
}
