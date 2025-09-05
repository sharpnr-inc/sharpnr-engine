package employeeservice

import (
	"context"
	"sharpnr-engine/core/models"

	"github.com/google/uuid"
)

func (s *EmployeeServiceImpl) UpdateEmployeeDetails(ctx context.Context, orgID uuid.UUID, emp *models.Employee) (*models.Employee, error) {
	return nil, nil
}
