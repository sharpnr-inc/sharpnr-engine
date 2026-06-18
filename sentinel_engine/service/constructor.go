package sentinelservice

import (
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinel_repo "sharpnr-engine/sentinel_engine/repo"
)

type sentinelService struct {
	repo      sentinel_repo.SentinelRepository
	txManager dbtx.TxManager
}

func NewSentinelService(
	repo sentinel_repo.SentinelRepository,
	txManager dbtx.TxManager,
) *sentinelService {
	return &sentinelService{
		repo:      repo,
		txManager: txManager,
	}
}
