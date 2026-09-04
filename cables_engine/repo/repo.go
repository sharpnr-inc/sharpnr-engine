package cablesEngineRepo

import (
	"context"
	"errors"

	cablesEngineErrs "sharpnr-engine/cables_engine/errs"
	cables_engine_models "sharpnr-engine/cables_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// conn returns the transaction the caller injected, or the repository's own
// handle when there is none, always bound to the request context so tracing and
// cancellation carry through.
func (c *cablesRepo) conn(ctx context.Context) *gorm.DB {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = c.db
	}
	return db.WithContext(ctx)
}

/* ---------------------------------- <***> --------------------------------- */

// IntializeOrgCablesAccount implements [CablesRepoInterface].
//
// The id is app-generated; created_at is left to GORM's convention hook.
func (c *cablesRepo) IntializeOrgCablesAccount(ctx context.Context, orgID uuid.UUID) (*cables_engine_models.OrgCablesAccount, error) {
	account := &cables_engine_models.OrgCablesAccount{
		ID:             uuid.New(),
		OrganizationID: orgID,
		IsActive:       true,
	}

	if err := c.conn(ctx).Create(account).Error; err != nil {
		return nil, err
	}

	return account, nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetCablesAccountDetails implements [CablesRepoInterface].
//
// "Never started" is a distinct answer rather than a bare gorm miss, because
// every other cables flow branches on exactly that.
func (c *cablesRepo) GetCablesAccountDetails(ctx context.Context, orgID uuid.UUID) (*cables_engine_models.OrgCablesAccount, error) {
	var account cables_engine_models.OrgCablesAccount

	err := c.conn(ctx).
		Where("organization_id = ?", orgID).
		First(&account).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cablesEngineErrs.ErrCablesServiceNotInitialized
		}
		return nil, err
	}

	return &account, nil
}

/* ---------------------------------- <***> --------------------------------- */

// CablesAccountSwitch implements [CablesRepoInterface].
//
// The account id is ANDed with the org rather than trusted on its own: an id
// belonging to another tenant affects no rows and reads back as not
// initialized.
func (c *cablesRepo) CablesAccountSwitch(ctx context.Context, orgID, orgCablesAccountID uuid.UUID, isActive bool) error {
	result := c.conn(ctx).
		Model(&cables_engine_models.OrgCablesAccount{}).
		Where("organization_id = ? AND id = ?", orgID, orgCablesAccountID).
		Update("is_active", isActive)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return cablesEngineErrs.ErrCablesServiceNotInitialized
	}
	return nil
}

/* ---------------------------------- <***> --------------------------------- */

// CreateChannel implements [CablesRepoInterface].
func (c *cablesRepo) CreateChannel(ctx context.Context, channel cables_engine_models.CablesChannel) error {
	return c.conn(ctx).Create(&channel).Error
}

/* ---------------------------------- <***> --------------------------------- */

// ListChannels implements [CablesRepoInterface].
//
// Ordered the way the sidebar renders: the explicit sort_order first, then
// name, so two channels sharing a position still come back in a stable order.
func (c *cablesRepo) ListChannels(ctx context.Context, orgID, orgCablesAccountID uuid.UUID) ([]cables_engine_models.CablesChannel, error) {
	var channels []cables_engine_models.CablesChannel

	err := c.conn(ctx).
		Where("organization_id = ? AND org_cables_account_id = ?", orgID, orgCablesAccountID).
		Order("sort_order ASC, name ASC").
		Find(&channels).
		Error

	if err != nil {
		return nil, err
	}

	return channels, nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetChannelByID implements [CablesRepoInterface].
//
// A channel belonging to another organization is not found here rather than
// returned — the tenant boundary is a predicate on the query.
func (c *cablesRepo) GetChannelByID(ctx context.Context, orgID, channelID uuid.UUID) (*cables_engine_models.CablesChannel, error) {
	var channel cables_engine_models.CablesChannel

	err := c.conn(ctx).
		Where("organization_id = ? AND id = ?", orgID, channelID).
		First(&channel).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cablesEngineErrs.ErrNoChannelFound
		}
		return nil, err
	}

	return &channel, nil
}

/* ---------------------------------- <***> --------------------------------- */

// ChannelSlugExists implements [CablesRepoInterface].
//
// GORM's soft-delete scope excludes deleted rows, so a deleted channel does not
// hold its slug hostage — the same rule the partial unique index enforces.
func (c *cablesRepo) ChannelSlugExists(ctx context.Context, orgCablesAccountID uuid.UUID, slug string) (bool, error) {
	if slug == "" {
		return false, nil
	}

	var count int64

	err := c.conn(ctx).
		Model(&cables_engine_models.CablesChannel{}).
		Where("org_cables_account_id = ? AND slug = ?", orgCablesAccountID, slug).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
