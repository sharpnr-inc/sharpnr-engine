package cortexEngineRepo

import (
	"context"

	dbtx "sharpnr-engine/pkg/dbTx"

	"gorm.io/gorm"
)

type cortexRepo struct {
	db *gorm.DB
}

func NewCortexEngineRepo(db *gorm.DB) *cortexRepo {
	return &cortexRepo{
		db: db,
	}
}

// conn returns the transaction the caller injected, or the repository's own
// handle when there is none, always bound to the request context so tracing and
// cancellation carry through.
func (c *cortexRepo) conn(ctx context.Context) *gorm.DB {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = c.db
	}
	return db.WithContext(ctx)
}
