package coreEmployeeServices

import (
	"context"
	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"

	"github.com/google/uuid"
)

func (s *employeeService) AddEmployeeInBulkInOrg(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, req []coreEmployeeContracts.AddEmployeeInBulkInOrganizationRequestDto) error {

	// in orchestrator we already have checked and verify all the emails
	// here we have all the checked emails first we will create the token and then save them to the
	// employee provision table and then we will send them email to set password


	return nil
}
