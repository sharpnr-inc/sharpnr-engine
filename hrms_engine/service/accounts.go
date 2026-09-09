package hrmsEngineService

import (
	"context"
	"errors"
	"time"

	hrmsEngineConstant "sharpnr-engine/hrms_engine/constant"
	hrmsEngineDtos "sharpnr-engine/hrms_engine/dtos"
	hrmsEngineErrs "sharpnr-engine/hrms_engine/errs"
	hrmsEngineModels "sharpnr-engine/hrms_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StartHrmsServiceForOrg implements [HrmsEngineServiceInterface].
//
// Provisioning is idempotent by intent rather than by accident: an org that
// already has an account never gets a second one (the unique index on
// organization_id would reject it anyway) — a switched-off account is turned
// back on, and a live one is told it is already live.
//
// The account and its settings row are written together. Every setting added
// later is then an ALTER plus an UPDATE, never a backfill hunting for orgs whose
// settings row was never created.
func (s *hrmsService) StartHrmsServiceForOrg(ctx context.Context, orgID, emplID uuid.UUID, req hrmsEngineDtos.StartHrmsServiceRequestDto) error {
	if !req.AgreeTermsAndConditions || !req.StartService {
		return hrmsEngineErrs.ErrTermsNotAccepted
	}

	existing, err := s.hrmsRepo.GetHrmsAccountByOrgID(ctx, orgID)
	if err != nil && !errors.Is(err, hrmsEngineErrs.ErrNoHRMSAccountFound) {
		return err
	}

	if existing != nil {
		if existing.IsActive {
			return hrmsEngineErrs.ErrHrmsServiceAlreadyStarted
		}

		return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
			txCtx := dbtx.InjectTx(ctx, tx)

			if err := s.hrmsRepo.HrmsAccountSwitch(txCtx, orgID, existing.ID, true); err != nil {
				return err
			}

			return s.recordAudit(txCtx, orgID, emplID, existing.ID,
				hrmsEngineConstant.AuditActionHrmsServiceResumed,
				"HRMS service was switched back on for the organization",
			)
		})
	}

	now := time.Now()

	account := hrmsEngineModels.OrgHrmsAccount{
		ID:             uuid.New(),
		OrganizationID: orgID,
		IsActive:       true,
		CreatedAt:      now,
	}

	settings := hrmsEngineModels.HrmsAccountSetting{
		ID:               uuid.New(),
		OrgHrmsAccountID: account.ID,
		OrganizationID:   orgID,
		UpdatedAt:        now,
	}

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.hrmsRepo.InitializeHrmsAccount(txCtx, account); err != nil {
			return err
		}

		if err := s.hrmsRepo.CreateSettingForHrmsAccount(txCtx, settings); err != nil {
			return err
		}

		return s.recordAudit(txCtx, orgID, emplID, account.ID,
			hrmsEngineConstant.AuditActionHrmsServiceStarted,
			"HRMS service was started for the organization",
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */

// GetHrmsAccountStatus implements [HrmsEngineServiceInterface].
//
// "Never started" is a zero-value answer rather than a 404: this is the read the
// client makes before it knows whether the org has HRMS at all, so a missing
// account is information, not a failure. A switched-off account still reports
// its id, because the client renders "turn it back on" against that account.
func (s *hrmsService) GetHrmsAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (hrmsEngineDtos.HrmsAccountStatusResponse, error) {
	account, err := s.hrmsRepo.GetHrmsAccountByOrgID(ctx, orgID)
	if err != nil {
		if errors.Is(err, hrmsEngineErrs.ErrNoHRMSAccountFound) {
			return hrmsEngineDtos.HrmsAccountStatusResponse{}, nil
		}
		return hrmsEngineDtos.HrmsAccountStatusResponse{}, err
	}

	return hrmsEngineDtos.HrmsAccountStatusResponse{
		ServiceStarted: true,
		IsActive:       &account.IsActive,
		ID:             &account.ID,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

// activeAccount is the gate every HRMS flow past provisioning goes through: it
// resolves the org's account and refuses a switched-off one.
//
// Not yet applied to ListOrgAllEmployees: that read is served by the core
// employees module, which is the org's roster and exists whether or not HRMS was
// ever started. The attendance, payroll and recruitment flows are HRMS's own
// data and go through here.
//
//lint:ignore U1000 gate for the HRMS flows landing on top of the account
func (s *hrmsService) activeAccount(ctx context.Context, orgID uuid.UUID) (*hrmsEngineModels.OrgHrmsAccount, error) {
	account, err := s.hrmsRepo.GetHrmsAccountByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	if !account.IsActive {
		return nil, hrmsEngineErrs.ErrAccountDisabled
	}

	return account, nil
}
