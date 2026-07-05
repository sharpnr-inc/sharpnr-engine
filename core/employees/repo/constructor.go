package employee_repo

import (
	"context"
	dbtx "sharpnr-engine/pkg/dbTx"

	"gorm.io/gorm"
)

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *employeeRepo {
	return &employeeRepo{
		db: db,
	}
}

// conn returns the transaction bound to the context if there is one, otherwise
// the plain DB. This is the same pattern used across the other repos so that
// writes join a caller's transaction transparently.
func (r *employeeRepo) conn(ctx context.Context) *gorm.DB {
	if tx, ok := dbtx.ExtractTx(ctx); ok {
		return tx
	}
	return r.db
}
