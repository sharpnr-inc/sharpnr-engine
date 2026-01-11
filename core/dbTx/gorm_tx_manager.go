package dbtx

import (
	"context"

	"gorm.io/gorm"
)

type GormTxManager struct {
	db *gorm.DB
}

func NewGormTxManager(db *gorm.DB) *GormTxManager {
	return &GormTxManager{db: db}
}

func (g *GormTxManager) WithTx(
	ctx context.Context,
	fn func(tx *gorm.DB) error,
) error {
	return g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
