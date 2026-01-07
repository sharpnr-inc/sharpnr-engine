package dbtx

import (
	"context"

	"gorm.io/gorm"
)

type TxManager interface {
	WithTx(ctx context.Context, fn func(tx *gorm.DB) error) error
}
