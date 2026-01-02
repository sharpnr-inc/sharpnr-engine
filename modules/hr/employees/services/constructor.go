package hr_employee_services

import (
	"context"
	hr_employee_repo "sharpnr-engine/modules/hr/employees/repo"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"
	sharphrdtos "sharpnr-engine/modules/hr/hr_dtos"

	"github.com/google/uuid"
)

type hrEmployeeService struct {
	hrEmpRepo hr_employee_repo.HREmployeeRepositoryInterface
}

// AddEmployeeInOrg implements HrEmployeeServiceInterface.
func (s *hrEmployeeService) AddEmployeeInOrg(ctx context.Context, sharpnr_id string, req sharphrdtos.CreateEmployeeInOrganization_RequestDTO) (*sharphrdtos.CreateEmployeeInOrganizationResponseDTO, error) {
	panic("unimplemented")
}

// RemoveEmployeeInOrg implements HrEmployeeServiceInterface.
func (s *hrEmployeeService) RemoveEmployeeInOrg(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) error {
	panic("unimplemented")
}

// UpdateEmployeeDetails implements HrEmployeeServiceInterface.
func (s *hrEmployeeService) UpdateEmployeeDetails(ctx context.Context, orgID uuid.UUID, emp *sharphr_models.Employee) (*sharphr_models.Employee, error) {
	panic("unimplemented")
}

func NewHrEmployeeService(hrEmpRepo hr_employee_repo.HREmployeeRepositoryInterface) *hrEmployeeService {
	return &hrEmployeeService{
		hrEmpRepo: hrEmpRepo,
	}
}
