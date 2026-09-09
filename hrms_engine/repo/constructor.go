package hrmsEngineRepo

import (
	"context"

	dbtx "sharpnr-engine/pkg/dbTx"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewHrmsEngineRepo(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

// conn returns the transaction the caller injected, or the repository's own
// handle when there is none, always bound to the request context so tracing and
// cancellation carry through.
func (d *repository) conn(ctx context.Context) *gorm.DB {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = d.db
	}
	return db.WithContext(ctx)
}
