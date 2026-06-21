package orchestrator_service

import (
	"context"

	"github.com/google/uuid"
)

// Fetches all the employee of the organization
func (s *orchestratorService) GetAllEmployeesOfOrg(ctx context.Context, orgID, emplID, propertyID uuid.UUID) (string, error) {

	return "", nil
}
