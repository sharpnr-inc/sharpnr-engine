package orchestrator_service

import (
	"context"
	"errors"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* ---------------------------------- <***> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                        Version 1.0.0 of Setup Wizard                       */
/* -------------------------------------------------------------------------- */

func (s *orchestratorService) SetupWiz_AddProperty(
	ctx context.Context,
	orgID uuid.UUID, emplID uuid.UUID,
	req *orchestrator_dtos.SetupOrgPropertyReqDTO,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		dbtx := dbtx.InjectTx(ctx, tx)

		stepParse_uuid, err := uuid.Parse(req.StepID)
		if err != nil {
			return err
		}

		input := org_setup_wizard_contracts.PropertySetupInput{
			StepID:       stepParse_uuid,
			OrgID:        orgID,
			Name:         req.Name,
			PropertyCode: *req.PropertyCode,
			AddressLine:  req.AddressLine,
			City:         req.City,
			State:        req.State,
			Country:      req.Country,
			Pincode:      req.Pincode,
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
func (s *orchestratorService) SetUpWizProgress(
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
func (s *orchestratorService) SetUpSuperAdminDetails(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, req *orchestrator_dtos.SetupSuperAdminDetailsReqDTO) error {
	// 1. Verify Org exists and is verified
	verified, err := s.organizationService.VerifyOrgByID(ctx, orgID)
	if err != nil {
		return err
	}

	if !verified {
		return errors.New("organization must be verified before setting up super admin details")
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

func (s *orchestratorService) SetupWiz_UpdateOrgName(ctx context.Context, orgID, emplID uuid.UUID, request *orchestrator_dtos.SetupWizUpdateOrgNameRequestDto) error {

	verified, err := s.organizationService.VerifyOrgByID(ctx, orgID)
	if err != nil {
		return err
	}

	if !verified {
		return errors.New("organization must be verified before setting up super admin details")
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
