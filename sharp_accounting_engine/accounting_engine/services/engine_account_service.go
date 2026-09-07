package accounting_service

import (
	"context"
	"errors"
	"time"

	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	auditEngineModels "sharpnr-engine/audit_log_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelContexts "sharpnr-engine/sentinel_contexts"
	accountingconstants "sharpnr-engine/sharp_accounting_engine/accounting_engine/constants"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"
	accountingEngineErrs "sharpnr-engine/sharp_accounting_engine/accounting_engine/errs"
	accounting_models "sharpnr-engine/sharp_accounting_engine/accounting_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StartAccountingServiceForOrg implements [AccountingServiceInterface].
//
// Provisioning is idempotent by intent rather than by accident: an org that
// already has an account never gets a second one (the unique index on
// organization_id would reject it anyway) — a switched-off account is turned
// back on, and a live one is told it is already live.
//
// The account and its settings row are written together. Every setting added
// later is then an ALTER plus an UPDATE, never a backfill hunting for orgs whose
// settings row was never created.
//
// This is the org-level switch only. It provisions no accounting book: a book
// belongs to a legal entity and is created by the legal-entity setup flow.
func (s *accountingServices) StartAccountingServiceForOrg(ctx context.Context, orgID, emplID uuid.UUID, req accountingdtos.StartAccountingServiceRequestDto) error {
	if !req.AgreeTermsAndConditions || !req.StartService {
		return accountingEngineErrs.ErrTermsNotAccepted
	}

	existing, err := s.accRepo.GetAccountingEngineAccountByOrgID(ctx, orgID)
	if err != nil && !errors.Is(err, accountingEngineErrs.ErrAccountingServiceNotInitialized) {
		return err
	}

	if existing != nil {
		if existing.IsActive {
			return accountingEngineErrs.ErrAccountingServiceAlreadyStarted
		}

		return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
			txCtx := dbtx.InjectTx(ctx, tx)

			if err := s.accRepo.AccountingEngineAccountSwitch(txCtx, orgID, existing.ID, true); err != nil {
				return err
			}

			return s.recordEngineAccountAudit(txCtx, orgID, emplID, existing.ID,
				accountingconstants.AuditActionAccountingServiceResumed,
				"Accounting service was switched back on for the organization",
			)
		})
	}

	now := time.Now()

	account := accounting_models.OrgAccountineEngineAccount{
		ID:             uuid.New(),
		OrganizationID: orgID,
		IsActive:       true,
		CreatedAt:      now,
	}

	settings := accounting_models.AccountingEngineAccountSetting{
		ID:                           uuid.New(),
		OrganizationID:               orgID,
		OrgAccountineEngineAccountID: account.ID,
		UpdatedAt:                    now,
	}

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.accRepo.InitializeAccountingEngineAccount(txCtx, account); err != nil {
			return err
		}

		if err := s.accRepo.CreateSettingForAccountingEngineAccount(txCtx, settings); err != nil {
			return err
		}

		return s.recordEngineAccountAudit(txCtx, orgID, emplID, account.ID,
			accountingconstants.AuditActionAccountingServiceStarted,
			"Accounting service was started for the organization",
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */

// GetAccountingAccountStatus implements [AccountingServiceInterface].
//
// "Never started" is a zero-value answer rather than a 404: this is the read the
// client makes before it knows whether the org has accounting at all, so a
// missing account is information, not a failure. A switched-off account still
// reports its id, because the client renders "turn it back on" against that
// account.
func (s *accountingServices) GetAccountingAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (accountingdtos.AccountingAccountStatusResponse, error) {
	account, err := s.accRepo.GetAccountingEngineAccountByOrgID(ctx, orgID)
	if err != nil {
		if errors.Is(err, accountingEngineErrs.ErrAccountingServiceNotInitialized) {
			return accountingdtos.AccountingAccountStatusResponse{}, nil
		}
		return accountingdtos.AccountingAccountStatusResponse{}, err
	}

	return accountingdtos.AccountingAccountStatusResponse{
		ServiceStarted: true,
		IsActive:       &account.IsActive,
		ID:             &account.ID,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

// activeEngineAccount is the gate the accounting-owned flows go through once
// the org-level switch is enforced.
//
// Deliberately not applied to the existing book, ledger and report reads yet:
// those orgs were provisioned before this account layer existed and have no row,
// so gating them now would break live accounting for every one of them. Apply it
// as those flows are revisited, once existing orgs have been backfilled.
//
//lint:ignore U1000 gate for the accounting flows once existing orgs are backfilled
func (s *accountingServices) activeEngineAccount(ctx context.Context, orgID uuid.UUID) (*accounting_models.OrgAccountineEngineAccount, error) {
	account, err := s.accRepo.GetAccountingEngineAccountByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	if !account.IsActive {
		return nil, accountingEngineErrs.ErrAccountingAccountIsDisabled
	}

	return account, nil
}

/* ---------------------------------- <***> --------------------------------- */

// recordEngineAccountAudit stamps one accounting-account event onto the org's
// history. It is called inside the same transaction as the write it describes,
// so the two commit or roll back together.
func (s *accountingServices) recordEngineAccountAudit(
	ctx context.Context,
	orgID, emplID, resourceID uuid.UUID,
	action, description string,
) error {
	if s.auditLogEngine == nil {
		return nil
	}

	// Protected routes have the client IP stamped onto ctx by the auth
	// middleware; when nothing injected it the field stays empty rather than
	// failing the write.
	var ipAddress string
	if authz, ok := sentinelContexts.Extract(ctx); ok {
		ipAddress = authz.Request.IPAddress
	}

	return s.auditLogEngine.Record(ctx, auditEngineDtos.RecordEventInput{
		OrganizationID: orgID,
		ActorID:        &emplID,
		ActorType:      auditEngineModels.ActorTypeEmployee,
		ActorIPAddress: ipAddress,
		Module:         accountingconstants.AuditModule,
		Action:         action,
		ResourceType:   accountingconstants.AuditResourceTypeAccountingAccount,
		ResourceID:     &resourceID,
		Description:    description,
	})
}
