package stockwaveEngineServices

import (
	"context"
	"errors"
	"time"

	dbtx "sharpnr-engine/pkg/dbTx"
	stockwaveEngineConstants "sharpnr-engine/stockwave_engine/constants"
	stockwaveEngineDtos "sharpnr-engine/stockwave_engine/dtos"
	stockwaveEngineErrs "sharpnr-engine/stockwave_engine/errs"
	stockwaveEngineModels "sharpnr-engine/stockwave_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StartStockwaveServiceForOrg implements [StockwaveEngineServicesInterface].
//
// Provisioning is idempotent by intent rather than by accident: an org that
// already has an account never gets a second one (the unique index on
// organization_id would reject it anyway) — a switched-off account is turned
// back on, and a live one is told it is already live.
//
// The account and its settings row are written together. Every setting added
// later is then an ALTER plus an UPDATE, never a backfill hunting for orgs whose
// settings row was never created.
func (s *stockwaveEngine) StartStockwaveServiceForOrg(ctx context.Context, orgID, emplID uuid.UUID, req stockwaveEngineDtos.StartStockwaveEngineRequestDto) error {
	if !req.AgreeTermsAndConditions || !req.StartService {
		return stockwaveEngineErrs.ErrTermsNotAccepted
	}

	existing, err := s.repo.GetStockwaveAccountByOrgID(ctx, orgID)
	if err != nil && !errors.Is(err, stockwaveEngineErrs.ErrNoAccountFound) {
		return err
	}

	if existing != nil {
		if existing.IsActive {
			return stockwaveEngineErrs.ErrStockwaveServiceAlreadyStarted
		}

		return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
			txCtx := dbtx.InjectTx(ctx, tx)

			if err := s.repo.StockwaveAccountSwitch(txCtx, orgID, existing.ID, true); err != nil {
				return err
			}

			return s.recordAudit(txCtx, orgID, emplID, existing.ID,
				stockwaveEngineConstants.AuditActionStockwaveServiceResumed,
				"Stockwave service was switched back on for the organization",
			)
		})
	}

	now := time.Now()

	account := stockwaveEngineModels.OrgStockwaveEngineAccounts{
		ID:             uuid.New(),
		OrganizationID: orgID,
		IsActive:       true,
		CreatedAt:      now,
	}

	settings := stockwaveEngineModels.StockwaveEngineAccountSetting{
		ID:                          uuid.New(),
		OrganizationID:              orgID,
		OrgStockwaveEngineAccountID: account.ID,
		UpdatedAt:                   now,
	}

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.repo.InitializeStockwaveAccount(txCtx, account); err != nil {
			return err
		}

		if err := s.repo.CreateSettingForStockwaveAccount(txCtx, settings); err != nil {
			return err
		}

		return s.recordAudit(txCtx, orgID, emplID, account.ID,
			stockwaveEngineConstants.AuditActionStockwaveServiceStarted,
			"Stockwave service was started for the organization",
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */

// GetStockwaveAccountStatus implements [StockwaveEngineServicesInterface].
//
// "Never started" is a zero-value answer rather than a 404: this is the read the
// client makes before it knows whether the org has stockwave at all, so a
// missing account is information, not a failure. A switched-off account still
// reports its id, because the client renders "turn it back on" against that
// account.
func (s *stockwaveEngine) GetStockwaveAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (stockwaveEngineDtos.StockwaveEngineStatusResponseDto, error) {
	account, err := s.repo.GetStockwaveAccountByOrgID(ctx, orgID)
	if err != nil {
		if errors.Is(err, stockwaveEngineErrs.ErrNoAccountFound) {
			return stockwaveEngineDtos.StockwaveEngineStatusResponseDto{}, nil
		}
		return stockwaveEngineDtos.StockwaveEngineStatusResponseDto{}, err
	}

	return stockwaveEngineDtos.StockwaveEngineStatusResponseDto{
		ServiceStarted: true,
		IsActive:       &account.IsActive,
		ID:             &account.ID,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

// activeAccount is the gate every stockwave flow past provisioning goes
// through: it resolves the org's account and refuses a switched-off one.
//
// Unused until the first rewritten stock flow lands — warehouses, items, stock
// movements — but it belongs here rather than being re-derived in each of them.
//
//lint:ignore U1000 gate for the stockwave flows landing on top of the account
func (s *stockwaveEngine) activeAccount(ctx context.Context, orgID uuid.UUID) (*stockwaveEngineModels.OrgStockwaveEngineAccounts, error) {
	account, err := s.repo.GetStockwaveAccountByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	if !account.IsActive {
		return nil, stockwaveEngineErrs.ErrAccountDisabled
	}

	return account, nil
}
