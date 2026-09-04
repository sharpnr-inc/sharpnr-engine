package orchestrator_service

import (
	"context"

	orgDepartmentsConstants "sharpnr-engine/core/departments/constants"
	orgDepartmentsDtos "sharpnr-engine/core/departments/dtos"
	orgDepartmentService "sharpnr-engine/core/departments/service"
	orchestrator_authz "sharpnr-engine/orchestrator/service/authz"

	"github.com/google/uuid"
)

// departmentOrchestrator implements OrgDepartmentOrchestratorServiceInterface.
//
// Every method is a single delegated call, so there is no transaction here
// (rule 2 in doc.go) — the department service owns its own consistency, and it
// already commits each write together with the audit entry describing it.
type departmentOrchestrator struct {
	guard                 *orchestrator_authz.Guard
	coreDepartmentService orgDepartmentService.OrgDepartmentServiceInterface
}

/* ---------------------------------- <***> --------------------------------- */

func (s *departmentOrchestrator) CreateNewOrgDeparment(ctx context.Context, orgID, emplID uuid.UUID, request orgDepartmentsDtos.CreateDepartmentRequestDto) error {
	// Type-level check: creating a department is not scoped to an instance,
	// because the instance does not exist yet.
	if err := s.guard.Authorize(ctx, orgID, emplID,
		orgDepartmentsConstants.ActionCanCreateDepartment,
		orgDepartmentsConstants.ResourceTypeDepartment, ""); err != nil {
		return err
	}

	return s.coreDepartmentService.CreateDepartment(ctx, orgID, emplID, request)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *departmentOrchestrator) ListAllOrgDepartments(ctx context.Context, orgID, emplID uuid.UUID) ([]orgDepartmentsDtos.DepartmentViewResponseDto, error) {
	if err := s.guard.Authorize(ctx, orgID, emplID,
		orgDepartmentsConstants.ActionCanViewDepartment,
		orgDepartmentsConstants.ResourceTypeDepartment, ""); err != nil {
		return nil, err
	}

	return s.coreDepartmentService.ListDepartments(ctx, orgID, emplID)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *departmentOrchestrator) GetOrgDepartmentDetails(ctx context.Context, orgID, emplID, departmentID uuid.UUID) (orgDepartmentsDtos.DepartmentViewResponseDto, error) {
	if err := s.guard.Authorize(ctx, orgID, emplID,
		orgDepartmentsConstants.ActionCanViewDepartment,
		orgDepartmentsConstants.ResourceTypeDepartment, departmentID.String()); err != nil {
		return orgDepartmentsDtos.DepartmentViewResponseDto{}, err
	}

	return s.coreDepartmentService.GetDepartmentByID(ctx, orgID, emplID, departmentID)
}

/* ---------------------------------- <***> --------------------------------- */

// UpdateOrgDepartment takes the id from the path rather than the body: a
// payload naming a different department must not be able to retarget the write
// past the instance-level check above it.
func (s *departmentOrchestrator) UpdateOrgDepartment(ctx context.Context, orgID, emplID, departmentID uuid.UUID, request orgDepartmentsDtos.UpdateDepartmentRequestDto) error {
	if err := s.guard.Authorize(ctx, orgID, emplID,
		orgDepartmentsConstants.ActionCanEditDepartment,
		orgDepartmentsConstants.ResourceTypeDepartment, departmentID.String()); err != nil {
		return err
	}

	request.ID = departmentID

	return s.coreDepartmentService.UpdateDepartment(ctx, orgID, emplID, request)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *departmentOrchestrator) DeleteOrgDepartment(ctx context.Context, orgID, emplID, departmentID uuid.UUID) error {
	if err := s.guard.Authorize(ctx, orgID, emplID,
		orgDepartmentsConstants.ActionCanDeleteDepartment,
		orgDepartmentsConstants.ResourceTypeDepartment, departmentID.String()); err != nil {
		return err
	}

	return s.coreDepartmentService.DeleteDepartment(ctx, orgID, emplID, departmentID)
}

/* ---------------------------------- <***> --------------------------------- */
