package folioEngineService

import (
	"context"
	"errors"
	"time"

	folioEngineConstants "sharpnr-engine/folio_engine/constants"
	folioEngineDtos "sharpnr-engine/folio_engine/dtos"
	folioEngineErrs "sharpnr-engine/folio_engine/errs"
	folioEngineModels "sharpnr-engine/folio_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StartFolioServiceForOrg implements [FolioEngineServiceInterface].
//
// Provisioning is idempotent by intent rather than by accident: an org that
// already has an account never gets a second one (the unique index on
// organization_id would reject it anyway) — a switched-off account is turned
// back on, and a live one is told it is already live.
//
// The account and its settings row are written together. Every setting added
// later is then an ALTER plus an UPDATE, never a backfill hunting for orgs whose
// settings row was never created.
func (f *folioService) StartFolioServiceForOrg(ctx context.Context, orgID, emplID uuid.UUID, req folioEngineDtos.StartFolioServiceRequestDto) error {
	if !req.AgreeTermsAndConditions || !req.StartService {
		return folioEngineErrs.ErrTermsNotAccepted
	}

	existing, err := f.folioRepo.GetFolioAccountByOrgID(ctx, orgID)
	if err != nil && !errors.Is(err, folioEngineErrs.ErrFolioServiceNotInitialized) {
		return err
	}

	if existing != nil {
		if existing.IsActive {
			return folioEngineErrs.ErrFolioServiceAlreadyStarted
		}

		return f.txManager.WithTx(ctx, func(tx *gorm.DB) error {
			txCtx := dbtx.InjectTx(ctx, tx)

			if err := f.folioRepo.FolioAccountSwitch(txCtx, orgID, existing.ID, true); err != nil {
				return err
			}

			return f.recordAudit(txCtx, orgID, emplID, existing.ID,
				folioEngineConstants.AuditActionFolioServiceResumed,
				"Folio service was switched back on for the organization",
			)
		})
	}

	now := time.Now()

	account := folioEngineModels.OrgFolioAccount{
		ID:             uuid.New(),
		OrganizationID: orgID,
		IsActive:       true,
		CreatedAt:      now,
	}

	settings := folioEngineModels.FolioAccountSetting{
		ID:                uuid.New(),
		OrgFolioAccountID: account.ID,
		OrganizationID:    orgID,
		UpdatedAt:         now,
	}

	return f.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := f.folioRepo.InitializeFolioAccount(txCtx, account); err != nil {
			return err
		}

		if err := f.folioRepo.CreateSettingForFolioAccount(txCtx, settings); err != nil {
			return err
		}

		return f.recordAudit(txCtx, orgID, emplID, account.ID,
			folioEngineConstants.AuditActionFolioServiceStarted,
			"Folio service was started for the organization",
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */

// GetFolioAccountStatus implements [FolioEngineServiceInterface].
//
// "Never started" is a zero-value answer rather than a 404: this is the read the
// client makes before it knows whether the org has Folio at all, so a missing
// account is information, not a failure. A switched-off account still reports
// its id, because the client renders "turn it back on" against that account.
func (f *folioService) GetFolioAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (folioEngineDtos.FolioAccountStatusResponse, error) {
	account, err := f.folioRepo.GetFolioAccountByOrgID(ctx, orgID)
	if err != nil {
		if errors.Is(err, folioEngineErrs.ErrFolioServiceNotInitialized) {
			return folioEngineDtos.FolioAccountStatusResponse{}, nil
		}
		return folioEngineDtos.FolioAccountStatusResponse{}, err
	}

	return folioEngineDtos.FolioAccountStatusResponse{
		ServiceStarted: true,
		IsActive:       &account.IsActive,
		ID:             &account.ID,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

// activeAccount is the gate every Folio flow past provisioning goes through: it
// resolves the org's account and refuses a switched-off one.
//
// Unused until the first real Folio flow lands — documents, sheets, boards —
// but it belongs here rather than being re-derived in each of them.
//
//lint:ignore U1000 gate for the Folio flows landing on top of the account
func (f *folioService) activeAccount(ctx context.Context, orgID uuid.UUID) (*folioEngineModels.OrgFolioAccount, error) {
	account, err := f.folioRepo.GetFolioAccountByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	if !account.IsActive {
		return nil, folioEngineErrs.ErrFolioAccountIsDisabled
	}

	return account, nil
}
