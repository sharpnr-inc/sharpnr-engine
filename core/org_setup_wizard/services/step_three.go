package org_setup_wizard_services

import (
	"context"
	orchestrator_dtos "sharpnr-engine/core/orchestrator/dtos"

	"github.com/google/uuid"
)

func (s *orgSetupWizardService) StepThreeLegalEntity(
	ctx context.Context,
	orgID uuid.UUID,
	req *orchestrator_dtos.SetupLegalEntityReqDTO,
) error {
	return nil

}
