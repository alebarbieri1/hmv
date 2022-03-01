package usecases

import "flavioltonon/hmv/domain/entity"

type AnalystUsecase interface {
	CreateAnalyst(user *entity.User) (*entity.Analyst, error)
	FindAnalystByID(analystID string) (*entity.Analyst, error)
}
