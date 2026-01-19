package org_setup_wizard_services

import (
	"context"

	"github.com/google/uuid"
)

// This si version V1 interface
// to Upgrade it Just add another interface making the
type OrgSetupWizardServiceInterface interface {
	CreateAndStartSetupStep(
		ctx context.Context,
		orgID uuid.UUID,
		currentStep string,
		status string,
	) error
}
