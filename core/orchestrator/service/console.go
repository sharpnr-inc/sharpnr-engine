package orchestrator_service

import (
	"context"

	"github.com/google/uuid"
)

/*
*This Method return the dto from the orchestrator,
*This method simply tell the if Organizations is Completed the setup or not
 */
func (o *orchestratorService) GetOrganizationSetupStatus(ctx context.Context, orgID uuid.UUID) (bool, error) {
	return false, nil
}
