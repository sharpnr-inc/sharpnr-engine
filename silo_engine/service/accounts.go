package siloEngineService

import (
	"context"
	"errors"
	"time"

	dbtx "sharpnr-engine/pkg/dbTx"
	siloEngineConstants "sharpnr-engine/silo_engine/constants"
	siloEngineDtos "sharpnr-engine/silo_engine/dtos"
	siloEngineErrs "sharpnr-engine/silo_engine/errs"
	siloEngineModels "sharpnr-engine/silo_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StartSiloServiceForOrg implements [SiloEngineServiceInterface].
//
// Provisioning is idempotent by intent rather than by accident: an org that
// already has an account never gets a second one (the unique index on
// organization_id would reject it anyway) — a switched-off account is turned
// back on, and a live one is told it is already live.
//
// The account and its settings row are written together. Every setting added
// later is then an ALTER plus an UPDATE, never a backfill hunting for orgs whose
// settings row was never created.
func (s *siloService) StartSiloServiceForOrg(ctx context.Context, orgID, emplID uuid.UUID, req siloEngineDtos.StartSiloServiceRequestDto) error {
	if !req.AgreeTermsAndConditions || !req.StartService {
		return siloEngineErrs.ErrTermsNotAccepted
	}

	existing, err := s.siloRepo.GetSiloAccountByOrgID(ctx, orgID)
	if err != nil && !errors.Is(err, siloEngineErrs.ErrSiloServiceNotInitialized) {
		return err
	}

	if existing != nil {
		if existing.IsActive {
			return siloEngineErrs.ErrSiloServiceAlreadyStarted
		}

		return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
			txCtx := dbtx.InjectTx(ctx, tx)

			if err := s.siloRepo.SiloAccountSwitch(txCtx, orgID, existing.ID, true); err != nil {
				return err
			}

			return s.recordAudit(txCtx, orgID, emplID, existing.ID,
				siloEngineConstants.AuditActionSiloServiceResumed,
				"Silo service was switched back on for the organization",
			)
		})
	}

	now := time.Now()

	account := siloEngineModels.OrgSiloAccount{
		ID:             uuid.New(),
		OrganizationID: orgID,
		IsActive:       true,
		CreatedAt:      now,
	}

	settings := siloEngineModels.SiloAccountSetting{
		ID:               uuid.New(),
		OrgSiloAccountID: account.ID,
		OrganizationID:   orgID,
		UpdtedAt:         now,
	}

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.siloRepo.InitializeSiloAccount(txCtx, account); err != nil {
			return err
		}

		if err := s.siloRepo.CreateSettingForSiloAccount(txCtx, settings); err != nil {
			return err
		}

		return s.recordAudit(txCtx, orgID, emplID, account.ID,
			siloEngineConstants.AuditActionSiloServiceStarted,
			"Silo service was started for the organization",
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */

// GetSiloAccountStatus implements [SiloEngineServiceInterface].
//
// "Never started" is a zero-value answer rather than a 404: this is the read the
// client makes before it knows whether the org has Silo at all, so a missing
// account is information, not a failure. A switched-off account still reports
// its id, because the client renders "turn it back on" against that account.
func (s *siloService) GetSiloAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (siloEngineDtos.SiloAccountStatusResponse, error) {
	account, err := s.siloRepo.GetSiloAccountByOrgID(ctx, orgID)
	if err != nil {
		if errors.Is(err, siloEngineErrs.ErrSiloServiceNotInitialized) {
			return siloEngineDtos.SiloAccountStatusResponse{}, nil
		}
		return siloEngineDtos.SiloAccountStatusResponse{}, err
	}

	return siloEngineDtos.SiloAccountStatusResponse{
		ServiceStarted: true,
		IsActive:       &account.IsActive,
		ID:             &account.ID,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

// activeAccount is the gate every Silo flow past provisioning goes through: it
// resolves the org's account and refuses a switched-off one.
//
// Unused until the first real Silo flow lands — buckets, databases, drive — but
// it belongs here rather than being re-derived in each of them.
//
//lint:ignore U1000 gate for the Silo flows landing on top of the account
func (s *siloService) activeAccount(ctx context.Context, orgID uuid.UUID) (*siloEngineModels.OrgSiloAccount, error) {
	account, err := s.siloRepo.GetSiloAccountByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	if !account.IsActive {
		return nil, siloEngineErrs.ErrSiloAccountIsDisabled
	}

	return account, nil
}
