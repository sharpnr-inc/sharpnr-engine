package sharpCrmEngineService

import (
	"context"
	"errors"
	"time"

	sharpCrmEngineConstants "sharpnr-engine/sharpCRM_engine/constants"
	sharpCrmEngineDtos "sharpnr-engine/sharpCRM_engine/dtos"
	sharpCrmEngineErrs "sharpnr-engine/sharpCRM_engine/errs"
	sharpCrmEngineModels "sharpnr-engine/sharpCRM_engine/models"

	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StartCrmServiceForOrg implements [SharpCrmEngineServiceInterface].
//
// Provisioning is idempotent by intent rather than by accident: an org that
// already has an account never gets a second one (the unique index on
// organization_id would reject it anyway) — a switched-off account is turned
// back on, and a live one is told it is already live.
//
// The account and its settings row are written together. Every setting added
// later is then an ALTER plus an UPDATE, never a backfill hunting for orgs whose
// settings row was never created.
func (r *Service) StartCrmServiceForOrg(ctx context.Context, orgID, emplID uuid.UUID, req sharpCrmEngineDtos.StartCRMServiceRequestDto) error {
	if !req.Agreed_terms_and_conditions || !req.StartService {
		return sharpCrmEngineErrs.ErrTermsNotAccepted
	}

	existing, err := r.crmRepo.GetCrmAccountByOrgID(ctx, orgID)
	if err != nil && !errors.Is(err, sharpCrmEngineErrs.ErrCrmServiceNotInitialized) {
		return err
	}

	if existing != nil {
		if existing.IsActive {
			return sharpCrmEngineErrs.ErrCrmServiceAlreadyStarted
		}

		return r.txManager.WithTx(ctx, func(tx *gorm.DB) error {
			txCtx := dbtx.InjectTx(ctx, tx)

			if err := r.crmRepo.CrmAccountSwitch(txCtx, orgID, existing.ID, true); err != nil {
				return err
			}

			return r.recordAudit(txCtx, orgID, emplID, existing.ID,
				sharpCrmEngineConstants.AuditActionCrmServiceResumed,
				"SharpCRM service was switched back on for the organization",
			)
		})
	}

	now := time.Now()

	account := sharpCrmEngineModels.OrgSharpCrmAccount{
		ID:             uuid.New(),
		OrganizationID: orgID,
		IsActive:       true,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	settings := sharpCrmEngineModels.SharpCrmAccountSetting{
		ID:                   uuid.New(),
		OrgSharpCrmAccountID: account.ID,
		OrganizationID:       orgID,
		LastUpdatedAt:        now,
	}

	return r.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := r.crmRepo.InitializeCrmAccount(txCtx, account); err != nil {
			return err
		}

		if err := r.crmRepo.CreateSettingForCrmAccount(txCtx, settings); err != nil {
			return err
		}

		return r.recordAudit(txCtx, orgID, emplID, account.ID,
			sharpCrmEngineConstants.AuditActionCrmServiceStarted,
			"SharpCRM service was started for the organization",
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */

// GetCrmAccountStatus implements [SharpCrmEngineServiceInterface].
//
// "Never started" is a zero-value answer rather than a 404: this is the read the
// client makes before it knows whether the org has SharpCRM at all, so a missing
// account is information, not a failure. A switched-off account still reports
// its id, because the client renders "turn it back on" against that account.
func (r *Service) GetCrmAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (sharpCrmEngineDtos.SharpCrmAccountStatusResponse, error) {
	account, err := r.crmRepo.GetCrmAccountByOrgID(ctx, orgID)
	if err != nil {
		if errors.Is(err, sharpCrmEngineErrs.ErrCrmServiceNotInitialized) {
			return sharpCrmEngineDtos.SharpCrmAccountStatusResponse{}, nil
		}
		return sharpCrmEngineDtos.SharpCrmAccountStatusResponse{}, err
	}

	return sharpCrmEngineDtos.SharpCrmAccountStatusResponse{
		ServiceStarted: true,
		IsActive:       &account.IsActive,
		ID:             &account.ID,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

// activeAccount is the gate every CRM flow past provisioning goes through: it
// resolves the org's account and refuses a switched-off one.
//
// Unused until the first real CRM flow lands — leads, deals, pipelines — but it
// belongs here rather than being re-derived in each of them.
//
//lint:ignore U1000 gate for the CRM flows landing on top of the account
func (r *Service) activeAccount(ctx context.Context, orgID uuid.UUID) (*sharpCrmEngineModels.OrgSharpCrmAccount, error) {
	account, err := r.crmRepo.GetCrmAccountByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	if !account.IsActive {
		return nil, sharpCrmEngineErrs.ErrCrmAccountIsDisabled
	}

	return account, nil
}
