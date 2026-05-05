package employee_services

import (
	"context"
	employee_contracts "sharpnr-engine/modules/hr/employees/contracts"

	"github.com/google/uuid"
)

func (s *employeeService) UpdateSuperAdminEmployee(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, updates *employee_contracts.UpdateSuperAdminEmployeeInput) error {
	panic("not implemented yet")
}
