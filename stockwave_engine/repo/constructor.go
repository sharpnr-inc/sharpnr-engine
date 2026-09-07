package stockwaveEngineRepo

import (
	"context"

	dbtx "sharpnr-engine/pkg/dbTx"

	"gorm.io/gorm"
)

type stockwaveRepo struct {
	db *gorm.DB
}

func NewStockwaveEngineRepo(db *gorm.DB) *stockwaveRepo {
	return &stockwaveRepo{
		db: db,
	}
}

// conn returns the transaction the caller injected, or the repository's own
// handle when there is none, always bound to the request context so tracing and
// cancellation carry through.
func (r *stockwaveRepo) conn(ctx context.Context) *gorm.DB {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	return db.WithContext(ctx)
}
