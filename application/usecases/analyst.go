package usecases

import "flavioltonon/hmv/domain/entity"

type AnalystUsecase interface {
	CreateAnalyst(userID string) (*entity.Analyst, error)
}
