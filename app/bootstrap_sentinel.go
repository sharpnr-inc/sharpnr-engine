package app

import (
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinel_repo "sharpnr-engine/sentinel/repo"
	sentinelservice "sharpnr-engine/sentinel/service"

	"gorm.io/gorm"
)

// bootstrapSentinel wires the Sentinel authoring service: repo over the DB,
// service over the repo + transaction manager. The engine is passed only as a
// cache invalidator so policy writes take effect immediately.
func bootstrapSentinel(
	db *gorm.DB,
	txManager dbtx.TxManager,
	engineCache sentinelservice.EngineCacheInvalidator,
) sentinelservice.SentinelServiceInterface {
	repo := sentinel_repo.NewSentinelRepo(db)
	return sentinelservice.NewSentinelService(repo, txManager, engineCache)
}
