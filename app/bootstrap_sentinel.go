package app

import (
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinel_repo "sharpnr-engine/sentinel_engine/repo"
	sentinelservice "sharpnr-engine/sentinel_engine/service"

	"gorm.io/gorm"
)

// bootstrapSentinel wires the Sentinel authoring service: repo over the DB,
// service over the repo + transaction manager. The authorization engine is not
// wired in here yet — policy/role/group CRUD is engine-independent.
func bootstrapSentinel(
	db *gorm.DB,
	txManager dbtx.TxManager,
) sentinelservice.SentinelService {
	repo := sentinel_repo.NewSentinelRepo(db)
	return sentinelservice.NewSentinelService(repo, txManager)
}
