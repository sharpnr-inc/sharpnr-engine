package app

import (
	sentinelengine "sharpnr-engine/sentinel_engine/engine"
	sentinel_repo "sharpnr-engine/sentinel_engine/repo"

	"gorm.io/gorm"
)

// bootstrapSentinelEngine constructs the authorization engine. It wraps the same
// Sentinel repo (which now also implements the engine's read interface) and is
// safe for concurrent use, so a single instance is shared across the app and
// injected into any service that needs to make authorization decisions.
func bootstrapSentinelEngine(db *gorm.DB) *sentinelengine.Sentinel {
	repo := sentinel_repo.NewSentinelRepo(db)
	return sentinelengine.New(repo)
}
