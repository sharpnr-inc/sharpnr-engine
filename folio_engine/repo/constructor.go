package folioEngineRepo

import (
	"context"

	dbtx "sharpnr-engine/pkg/dbTx"

	"gorm.io/gorm"
)

type folioRepo struct {
	db *gorm.DB
}

func NewFolioEngineRepo(db *gorm.DB) *folioRepo {
	return &folioRepo{
		db: db,
	}
}

// conn returns the transaction the caller injected, or the repository's own
// handle when there is none, always bound to the request context so tracing and
// cancellation carry through.
func (f *folioRepo) conn(ctx context.Context) *gorm.DB {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = f.db
	}
	return db.WithContext(ctx)
}
