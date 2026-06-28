package sentinel_repo

import (
	"context"

	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type sentinelRepo struct {
	db *gorm.DB
}

// CreateUserRole implements [SentinelRepositoryInterface].
func (r *sentinelRepo) CreateUserRole(ctx context.Context, userRole *sentinelengine_models.UserRole) error {
	panic("unimplemented")
}

// UpdateRoleDetails implements [SentinelRepositoryInterface].
func (r *sentinelRepo) UpdateRoleDetails(ctx context.Context, roleID uuid.UUID, updates *sentinelDtos.UpdateRoleDetailsRequestDto) error {
	panic("unimplemented")
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
