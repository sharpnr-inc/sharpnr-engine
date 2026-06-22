package org_signup_services

import (
	"context"
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *orgSignUpService) VerifyOtp(
	ctx context.Context,
	otp string,
	slug uuid.UUID,
) (orgID uuid.UUID, emp uuid.UUID, err error) {

	// The whole flow promotes a pending signup into a live org: it creates the
	// organization, its founding super-admin employee, the authorization
	// defaults (owner role + policy), and the setup-wizard progress. These must
	// be atomic — a half-provisioned org (e.g. an employee with no owner role,
	// or an org with no wizard progress) would be unrecoverable. Run them in one
	// transaction; the inner services join it via the tx-aware context, so
	// ProvisionOrganization's own WithTx reuses this transaction rather than
	// opening a nested one.
	err = r.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		dataForEmp, err := r.organizationService.VerifySlugCreateDefaultOrg(txCtx, otp, slug)
		if err != nil {
			return err
		}

		output, err := r.organizationService.CreateNewOrganization(txCtx, dataForEmp)
		if err != nil {
			return err
		}

		empId, err := r.employeeSvc.CreateDefaultSuperAdmin(txCtx, output.OrgID, output.Email, output.HashedPass)
		if err != nil {
			return err
		}

		// Authorization: give the founding employee a full-access "Owner" role
		if err := r.sentinelService.ProvisionOrganization(txCtx, output.OrgID, empId); err != nil {
			return err
		}

		// * This is Add Progress in Setup Wizard
		// * and fetch the setup progress it will put in the org_setup_wizard table and then it will be show in the UI that the user is in the Super Admin Profile step
		if err := r.setUpWizService.AddStepProgress(
			txCtx,
			output.OrgID,
			org_setup_wizards_constants.SetupStepSuperAdminProfile,
			org_setup_wizards_constants.SetupStatusInProgress); err != nil {
			return err
		}

		orgID, emp = output.OrgID, empId
		return nil
	})
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	return orgID, emp, nil
}
