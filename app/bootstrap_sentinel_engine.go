package app

import (
	sentinel_repo "sharpnr-engine/sentinel/repo"
	sentinelEngine "sharpnr-engine/sentinel_engine"

	"gorm.io/gorm"
)

// bootstrapSentinelEngine constructs the authorization engine. It wraps the same
// Sentinel repo (which also implements the engine's read interface) and is safe
// for concurrent use, so a single instance is shared across the app and
// injected into any service that needs to make authorization decisions.
func bootstrapSentinelEngine(db *gorm.DB) *sentinelEngine.SentinelEngine {
	return sentinelEngine.NewSentinelEngine(sentinel_repo.NewSentinelRepo(db))
}
