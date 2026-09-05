package siloEngineRepo

import (
	"context"

	dbtx "sharpnr-engine/pkg/dbTx"

	"gorm.io/gorm"
)

type siloRepo struct {
	db *gorm.DB
}

func NewSiloEngineRepo(db *gorm.DB) *siloRepo {
	return &siloRepo{
		db: db,
	}
}

// conn returns the transaction the caller injected, or the repository's own
// handle when there is none, always bound to the request context so tracing and
// cancellation carry through.
func (s *siloRepo) conn(ctx context.Context) *gorm.DB {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = s.db
	}
	return db.WithContext(ctx)
}
