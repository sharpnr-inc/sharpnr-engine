package orchestrator_service

import (
	"context"

	"github.com/google/uuid"
)

/*
*This Method return the dto from the orchestrator,
*This method simply tell the if Organizations is Completed the setup or not
 */
func (o *orchestratorService) GetOrganizationStatus(ctx context.Context, orgID uuid.UUID) (bool, error) {

	status, err := o.organizationService.GetStatus(ctx, orgID)
	if err != nil {
		return false, err
	}

	return status != "NOT_CONFIGURED", nil
}
