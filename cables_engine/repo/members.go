package cablesEngineRepo

import (
	"context"
	"errors"
	"time"

	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	cablesEngineErrs "sharpnr-engine/cables_engine/errs"
	cables_engine_models "sharpnr-engine/cables_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UpsertEmployeeMembers implements [CablesRepoInterface].
//
// One statement for the whole roster rather than a read-then-write per
// employee: the unique index on (account, employee_id) turns the sync into an
// upsert, and DoUpdates refreshes only the display name — presence, custom
// status and the member's own id must survive a sync, or every roster refresh
// would knock everyone offline and orphan the ids messages point at.
func (c *cablesRepo) UpsertEmployeeMembers(ctx context.Context, members []cables_engine_models.CablesMember) error {
	if len(members) == 0 {
		return nil
	}

	return c.conn(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "org_cables_account_id"}, {Name: "employee_id"}},
			// The index is partial, so Postgres can only infer it when the
			// conflict target repeats its predicate — without this the
			// statement fails with "no unique or exclusion constraint matching
			// the ON CONFLICT specification".
			TargetWhere: clause.Where{Exprs: []clause.Expression{
				clause.Expr{SQL: "employee_id IS NOT NULL AND deleted_at IS NULL"},
			}},
			DoUpdates: clause.AssignmentColumns([]string{"display_name", "updated_at"}),
		}).
		Create(&members).
		Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetMemberByEmployee implements [CablesRepoInterface].
func (c *cablesRepo) GetMemberByEmployee(ctx context.Context, orgCablesAccountID, employeeID uuid.UUID) (*cables_engine_models.CablesMember, error) {
	var member cables_engine_models.CablesMember

	err := c.conn(ctx).
		Where("org_cables_account_id = ? AND employee_id = ?", orgCablesAccountID, employeeID).
		First(&member).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cablesEngineErrs.ErrMemberNotFound
		}
		return nil, err
	}

	return &member, nil
}

/* ---------------------------------- <***> --------------------------------- */

// ListMembers implements [CablesRepoInterface]: the account's roster, which the
// client caches and renders every author and mention from.
//
// Guests are excluded: an ephemeral identity minted for one meeting has no
// place in the org directory.
func (c *cablesRepo) ListMembers(ctx context.Context, orgCablesAccountID uuid.UUID) ([]cables_engine_models.CablesMember, error) {
	var members []cables_engine_models.CablesMember

	err := c.conn(ctx).
		Where("org_cables_account_id = ? AND is_ephemeral = false", orgCablesAccountID).
		Order("display_name ASC").
		Find(&members).
		Error

	return members, err
}

/* ---------------------------------- <***> --------------------------------- */

// SetMemberPresence implements [CablesRepoInterface]. Driven by websocket
// connect/disconnect, never by a CRUD handler.
func (c *cablesRepo) SetMemberPresence(ctx context.Context, memberID uuid.UUID, presence cablesEngineConstants.PresenceStatus) error {
	now := time.Now()

	return c.conn(ctx).
		Model(&cables_engine_models.CablesMember{}).
		Where("id = ?", memberID).
		Updates(map[string]any{
			"presence":     presence,
			"last_seen_at": now,
			"updated_at":   now,
		}).
		Error
}
