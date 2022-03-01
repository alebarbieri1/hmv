package repositories

import "flavioltonon/hmv/domain/entity"

type AnalystsRepository interface {
	CreateAnalyst(pacient *entity.Analyst) error
	FindAnalystByID(analysisID string) (*entity.Analyst, error)
	FindAnalystByUserID(userID string) (*entity.Analyst, error)
}
