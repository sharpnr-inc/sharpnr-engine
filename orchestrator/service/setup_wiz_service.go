package orchestrator_service

import (
	"context"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
	orgSetUpWizardErrs "sharpnr-engine/core/org_setup_wizard/errs"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	organization_services "sharpnr-engine/core/organization/services"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// setupWizardOrchestrator implements OrchestratorSetupWizardInterface: the
// step-by-step onboarding flow that runs before the org is fully configured.
type setupWizardOrchestrator struct {
	txManager           dbtx.TxManager
	orgSetupWizService  org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface
	organizationService organization_services.OrganizationServiceInterface
}

/* ---------------------------------- <***> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                        Version 1.0.0 of Setup Wizard                       */
/* -------------------------------------------------------------------------- */

// derefOrEmpty returns the value behind an optional request field, or the empty
// string when the client omitted it.
func derefOrEmpty(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func (s *setupWizardOrchestrator) SetupWiz_AddProperty(
	ctx context.Context,
	orgID uuid.UUID, emplID uuid.UUID,
	req *orchestrator_dtos.SetupOrgPropertyReqDTO,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		dbtx := dbtx.InjectTx(ctx, tx)

		stepParse_uuid, err := uuid.Parse(req.StepID)
		if err != nil {
			return orgSetUpWizardErrs.InvalidStepIDErr
		}

		if err := req.Validate(); err != nil {
			return err
		}

		// Every property field is optional on the wire so the step can be skipped,
		// so they are dereferenced defensively rather than assumed present.
		input := org_setup_wizard_contracts.PropertySetupInput{
			StepID:       stepParse_uuid,
			Skipped:      req.Skipped,
			OrgID:        orgID,
			Name:         derefOrEmpty(req.Name),
			PropertyCode: derefOrEmpty(req.PropertyCode),
			AddressLine:  derefOrEmpty(req.AddressLine),
			City:         derefOrEmpty(req.City),
			State:        derefOrEmpty(req.State),
			Country:      derefOrEmpty(req.Country),
			Pincode:      derefOrEmpty(req.Pincode),
		}

		if err := s.orgSetupWizService.Step_two_AddProperty(dbtx, orgID, emplID, input); err != nil {
			return err
		}

		return nil

	})
}

/* ---------------------------------- <***> --------------------------------- */

// This method retrieves the current progress of the setup wizard for a given organization.
// It returns a structured response containing the overall status, current step, progress percentage, completion
// status, and details of each step in the wizard.
func (s *setupWizardOrchestrator) SetUpWizProgress(
	ctx context.Context,
	orgID uuid.UUID,
) (orchestrator_dtos.SetupProgressResponseDTO, error) {

	progress, err := s.orgSetupWizService.Progress(ctx, orgID)
	if err != nil {
		return orchestrator_dtos.SetupProgressResponseDTO{}, err
	}

	steps := make([]orchestrator_dtos.SetupStepProgressResponseDTO, 0, len(progress.Steps))

	for _, step := range progress.Steps {
		steps = append(steps, orchestrator_dtos.SetupStepProgressResponseDTO{
			ID:          step.ID,
			Step:        step.Step,
			Status:      step.Status,
			Version:     step.Version,
			StartedAt:   step.StartedAt,
			CompletedAt: step.CompletedAt,
		})
	}

	return orchestrator_dtos.SetupProgressResponseDTO{
		OverallStatus: progress.OverallStatus,
		CurrentStep:   progress.CurrentStep,
		IsCompleted:   progress.IsCompleted,
		Steps:         steps,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

// This super admin completed details will be their regardless of the version
// this is method completes the super admin details
//
// it takes organizationID and emplID id that we created after the signup process
func (s *setupWizardOrchestrator) SetUpSuperAdminDetails(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, req *orchestrator_dtos.SetupSuperAdminDetailsReqDTO) error {
	// 1. Verify Org exists and is verified
	verified, err := s.organizationService.VerifyOrgByID(ctx, orgID)
	if err != nil {
		return err
	}

	if !verified {
		return orgSetUpWizardErrs.OrgNotVerifiedErr
	}

	// 2. Now we will call the identity service to update the super admin details
	if err := s.orgSetupWizService.UpdateSuperAdminDetailsStep(ctx, orgID, emplID, org_setup_wizard_contracts.SuperAdminDetailsInput{
		StepID:         req.StepID,
		OrgID:          orgID,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		SecondaryEmail: req.SecondaryEmail,
		SecondaryPhone: "",
		EmployeeID:     *req.EmployeeID,
	}); err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *setupWizardOrchestrator) SetupWiz_UpdateOrgName(ctx context.Context, orgID, emplID uuid.UUID, request *orchestrator_dtos.SetupWizUpdateOrgNameRequestDto) error {

	verified, err := s.organizationService.VerifyOrgByID(ctx, orgID)
	if err != nil {
		return err
	}

	if !verified {
		return orgSetUpWizardErrs.OrgNotVerifiedErr
	}

	if err := s.orgSetupWizService.UpdateOrgDetailsStep(ctx, orgID, emplID, org_setup_wizard_contracts.OrganizationDetailsSetupInput{
		StepID: request.StepID,
		OrgID:  orgID,
		Name:   request.OrgName,
	}); err != nil {
		return err
	}

	return nil

}
