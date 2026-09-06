package cortexEngineService

import (
	"context"
	"errors"
	"time"

	cortexEngineConstants "sharpnr-engine/cortex_engine/constants"
	cortexEngineDtos "sharpnr-engine/cortex_engine/dtos"
	cortexEngineErrs "sharpnr-engine/cortex_engine/errs"
	cortexEngineModels "sharpnr-engine/cortex_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StartCortexServiceForOrg implements [CortexEngineServiceInterface].
//
// Provisioning is idempotent by intent rather than by accident: an org that
// already has an account never gets a second one (the unique index on
// organization_id would reject it anyway) — a switched-off account is turned
// back on, and a live one is told it is already live.
//
// The account and its settings row are written together. Every setting added
// later is then an ALTER plus an UPDATE, never a backfill hunting for orgs whose
// settings row was never created.
func (c *coretexService) StartCortexServiceForOrg(ctx context.Context, orgID, emplID uuid.UUID, req cortexEngineDtos.StartCortexServiceRequestDto) error {
	if !req.AgreeTermsAndCondiotions || !req.StartService {
		return cortexEngineErrs.ErrTermsNotAccepted
	}

	existing, err := c.cortexRepo.GetCortexAccountByOrgID(ctx, orgID)
	if err != nil && !errors.Is(err, cortexEngineErrs.ErrNoCortexAccountFound) {
		return err
	}

	if existing != nil {
		if existing.IsActive {
			return cortexEngineErrs.ErrCortexServiceAlreadyStarted
		}

		return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
			txCtx := dbtx.InjectTx(ctx, tx)

			if err := c.cortexRepo.CortexAccountSwitch(txCtx, orgID, existing.ID, true); err != nil {
				return err
			}

			return c.recordAudit(txCtx, orgID, emplID, existing.ID,
				cortexEngineConstants.AuditActionCortexServiceResumed,
				"Cortex service was switched back on for the organization",
			)
		})
	}

	now := time.Now()

	account := cortexEngineModels.OrgCortexAccount{
		ID:             uuid.New(),
		OrganizationID: orgID,
		IsActive:       true,
		CreatedAt:      now,
	}

	settings := cortexEngineModels.CortexAccountSetting{
		ID:                 uuid.New(),
		OrgCortexAccountID: account.ID,
		OrganizationID:     orgID,
		UpdatedAt:          now,
	}

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := c.cortexRepo.InitializeCortexAccount(txCtx, account); err != nil {
			return err
		}

		if err := c.cortexRepo.CreateSettingForCortexAccount(txCtx, settings); err != nil {
			return err
		}

		return c.recordAudit(txCtx, orgID, emplID, account.ID,
			cortexEngineConstants.AuditActionCortexServiceStarted,
			"Cortex service was started for the organization",
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */

// GetCortexAccountStatus implements [CortexEngineServiceInterface].
//
// "Never started" is a zero-value answer rather than a 404: this is the read the
// client makes before it knows whether the org has Cortex at all, so a missing
// account is information, not a failure. A switched-off account still reports
// its id, because the client renders "turn it back on" against that account.
func (c *coretexService) GetCortexAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (cortexEngineDtos.CortexAccountStatusResponse, error) {
	account, err := c.cortexRepo.GetCortexAccountByOrgID(ctx, orgID)
	if err != nil {
		if errors.Is(err, cortexEngineErrs.ErrNoCortexAccountFound) {
			return cortexEngineDtos.CortexAccountStatusResponse{}, nil
		}
		return cortexEngineDtos.CortexAccountStatusResponse{}, err
	}

	return cortexEngineDtos.CortexAccountStatusResponse{
		ServiceStarted: true,
		IsActive:       &account.IsActive,
		ID:             &account.ID,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

// activeAccount is the gate every Cortex flow past provisioning goes through:
// it resolves the org's account and refuses a switched-off one.
//
// Unused until the first real Cortex flow lands — roles, policies, API keys —
// but it belongs here rather than being re-derived in each of them.
//
//lint:ignore U1000 gate for the Cortex flows landing on top of the account
func (c *coretexService) activeAccount(ctx context.Context, orgID uuid.UUID) (*cortexEngineModels.OrgCortexAccount, error) {
	account, err := c.cortexRepo.GetCortexAccountByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	if !account.IsActive {
		return nil, cortexEngineErrs.ErrCortexAccountIsDisabled
	}

	return account, nil
}
