package sentinel_repo

import (
	"context"

	dbtx "sharpnr-engine/pkg/dbTx"

	"gorm.io/gorm"
)

// Compile-time guard: the concrete repo must satisfy the full interface.
var _ SentinelRepositoryInterface = (*sentinelRepo)(nil)

type sentinelRepo struct {
	db *gorm.DB
}

func NewSentinelRepo(db *gorm.DB) *sentinelRepo {
	return &sentinelRepo{db: db}
}

// conn returns the transaction bound to the context if there is one, otherwise
// the plain DB. This is the same pattern used across the other repos so that
// writes join a caller's transaction transparently.
func (r *sentinelRepo) conn(ctx context.Context) *gorm.DB {
	if tx, ok := dbtx.ExtractTx(ctx); ok {
		return tx
	}
	return r.db
}
