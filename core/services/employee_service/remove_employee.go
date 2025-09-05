package employeeservice

import (
	"context"

	"github.com/google/uuid"
)

func (s *EmployeeServiceImpl) RemoveEmployeeInOrg(ctx context.Context, orgID, empID uuid.UUID) error {
	return nil
}
