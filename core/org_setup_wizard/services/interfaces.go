package org_setup_wizard_services

import (
	"context"

	"github.com/google/uuid"
)

type OrgSetupWizardServiceInterface interface {
	CreateAndStartSetupStep(
		ctx context.Context,
		orgID uuid.UUID,
		currentStep string,
		status string,
	) error
}
