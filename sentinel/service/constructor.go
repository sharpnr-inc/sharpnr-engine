package sentinelservice

import (
	"context"

	dbtx "sharpnr-engine/pkg/dbTx"
	sentinel_repo "sharpnr-engine/sentinel/repo"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// EngineCacheInvalidator is the one engine capability the authoring service
// needs: dropping an org's compiled policy cache after a write. Declared here
// (instead of importing the engine package) so the service stays decoupled from
// evaluation; *sentinelEngine.SentinelEngine satisfies it.
type EngineCacheInvalidator interface {
	InvalidateOrganization(orgID uuid.UUID)
}

type sentinelService struct {
	repo        sentinel_repo.SentinelRepositoryInterface
	txManager   dbtx.TxManager
	engineCache EngineCacheInvalidator
}

func NewSentinelService(
	repo sentinel_repo.SentinelRepositoryInterface,
	txManager dbtx.TxManager,
	engineCache EngineCacheInvalidator,
) *sentinelService {
	return &sentinelService{
		repo:        repo,
		txManager:   txManager,
		engineCache: engineCache,
	}
}

// withTx runs fn inside one DB transaction, with the tx injected into the
// context so every repo call joins it (the codebase-wide dbtx pattern).
func (s *sentinelService) withTx(ctx context.Context, fn func(txCtx context.Context) error) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		return fn(dbtx.InjectTx(ctx, tx))
	})
}

// invalidateEngine drops the org's compiled policies after a policy write. Nil
// engineCache (engine not wired, e.g. in tests) is a no-op.
func (s *sentinelService) invalidateEngine(orgID uuid.UUID) {
	if s.engineCache != nil {
		s.engineCache.InvalidateOrganization(orgID)
	}
}
