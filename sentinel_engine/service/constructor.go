package sentinelservice

import (
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelengine "sharpnr-engine/sentinel_engine/engine"
	sentinel_repo "sharpnr-engine/sentinel_engine/repo"
)

type sentinelService struct {
	repo      sentinel_repo.SentinelRepositoryInterface
	txManager dbtx.TxManager
	// engine is needed so policy-mutating operations can invalidate the org's
	// compiled PolicySet cache after regenerating its Cedar.
	engine *sentinelengine.Sentinel
}

func NewSentinelService(
	repo sentinel_repo.SentinelRepositoryInterface,
	txManager dbtx.TxManager,
	engine *sentinelengine.Sentinel,
) *sentinelService {
	return &sentinelService{
		repo:      repo,
		txManager: txManager,
		engine:    engine,
	}
}
