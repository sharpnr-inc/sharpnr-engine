package coreEmployeeServices

import (
	"context"

	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"

	"github.com/google/uuid"
)

type EmployeeGeneralServiceInterface interface {
	IsUserExists(ctx context.Context, employeeEmail string) (bool, error)
}

/* ---------------------------------- <***> --------------------------------- */

type EmployeeSignUpServiceInterface interface {
	EmployeeGeneralServiceInterface
	CreateDefaultSuperAdmin(ctx context.Context, orgId uuid.UUID, rootEmail string, hasPass string) (uuid.UUID, error)
}

/* ---------------------------------- <***> --------------------------------- */

// This interface is used by authentication modules Employee will directly interact with the employee module through this interface.
// This interface will have employee Ops which are used by the other module interact with the employee module.
type EmployeeAuthenticationServiceInterface interface {
	EmployeeGeneralServiceInterface
	AuthenticateEmployee(ctx context.Context, email string, pass string) (*coreEmployeeContracts.EmployeeReader, error)
	GetEmployeeProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) (*coreEmployeeContracts.EmployeeProfileOutput, error)
	GetEmployeeDetails(ctx context.Context, orgID, emplID uuid.UUID) (*coreEmployeeContracts.EmployeeDetailsOutput, error)

	// Mobile refresh-token sessions. See mobile_session_service.go for the
	// rotation and reuse-detection rules.
	StartMobileSession(ctx context.Context, employeeID, orgID uuid.UUID, role string) (sessionID, jti uuid.UUID, err error)
	RotateMobileSession(ctx context.Context, sessionID, presentedJTI uuid.UUID) (newJTI uuid.UUID, role string, err error)
	EndMobileSession(ctx context.Context, sessionID uuid.UUID) error
	EndAllMobileSessions(ctx context.Context, employeeID uuid.UUID) (int64, error)
	PurgeExpiredMobileSessions(ctx context.Context) (int64, error)
}

/* ---------------------------------- <***> --------------------------------- */

// this is interface which will used by the HR module to interact with the employee module, this interface the have employee Ops used by the Org Hr
type EmployeeHRServiceInterface interface {
	// General service which will used by the other module to interact with the employee module,
	EmployeeGeneralServiceInterface
	ListOrgAllEmployees(ctx context.Context, orgID, emplID uuid.UUID) ([]*coreEmployeeContracts.OrgEmployeeList, error)
	// Add Employees in bulk in the organization
	AddEmployeeInBulkInOrg(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, req []coreEmployeeContracts.AddEmployeeInBulkInOrganizationRequestDto) error
}

/* ---------------------------------- <***> --------------------------------- */

type EmployeeIamEngineServiceInterface interface {
	// this is General service which will used by the other module to interact with the employee module,
	EmployeeGeneralServiceInterface
	GetEmployeeProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) (*coreEmployeeContracts.EmployeeProfileOutput, error)

	// ListOrgEmployees is the org roster as plain identity rows: id, name,
	// email, status. Deliberately thinner than the HR listing, which decorates
	// every employee with a sentinel role lookup — a consumer that only needs
	// names should not pay for that, and must not fail when an employee has no
	// role. Cables reads this once per roster sync, never on a hot path.
	ListOrgEmployees(ctx context.Context, orgID uuid.UUID) ([]coreEmployeeContracts.OrgEmployeeIdentity, error)
	// GetEmployeeEmploymentDetails()
}

type EmployeeOrchestratorServiceInterface interface {
	// since only orchestrator can create a new superAdmin/Owner and only a superAdmin/Owner can add a owner
	//
	// emplID is the acting owner; the returned id is the employee that was
	// created. Whether the caller is allowed to do this is the orchestrator's
	// call — see AddNewSuperAdmin there.
	AddNewOwnerOrSuperAdmin(ctx context.Context, orgID, emplID uuid.UUID, req coreEmployeeContracts.AddNewSuperAdminInOrganizationRequestDto) (uuid.UUID, error)
}

type EmployeeSetupWizServiceInterface interface {
	UpdateSuperAdminEmployee(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, updates *coreEmployeeContracts.UpdateSuperAdminEmployeeInput) error
}
