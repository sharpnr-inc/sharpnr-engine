package org_setup_wizard_services

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *orgSetupWizardService) Step_two_AddProperty(ctx context.Context, orgID, emplID uuid.UUID, input org_setup_wizard_contracts.PropertySetupInput) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		dbTx := dbtx.InjectTx(ctx, tx)

		stepstatus, err := s.StepVerication(dbTx, orgID, emplID, input.StepID)
		if err != nil {
			return err
		}

		if *stepstatus != org_setup_wizards_constants.SetupStatusInProgress {
			return ErrorInvalidStepStatus
		}

		mappedData := org_property_contracts.CreateOrgPropertyInput{
			OrgID:       input.OrgID,
			Name:        input.Name,
			AddressLine: input.AddressLine,
			City:        input.City,
			State:       input.State,
			Country:     input.Country,
			Pincode:     input.Pincode,
		}

		if err := s.orgPropertySvc.CreatePropertyDefault(dbTx, mappedData); err != nil {
			return err
		}

		if err := s.orgSetWizRepo.UpdateCurrentStepStatus(
			dbTx, org_setup_wizards_constants.SetupStatusCompleted,
			input.OrgID, input.StepID,
		); err != nil {
			return err
		}

		return nil

	})

}
