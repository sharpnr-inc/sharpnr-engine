package orchestrator_repo

import "gorm.io/gorm"

type orchestratorRepo struct {
	db *gorm.DB
}

func NewOrchestratorRepository(db *gorm.DB) *orchestratorRepo {
	return &orchestratorRepo{
		db: db,
	}
}
