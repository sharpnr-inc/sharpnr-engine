package coreEmployeeRepo

import (
	"context"
	coreEmployeeModels "sharpnr-engine/core/employees/models"

	"time"

	"github.com/google/uuid"
)

type EmployeeRepositoryInterface interface {
	EmployeeMainRepositoryInterface
	EmployeeOnboardingRepositoryInterface
	EmployeeAuthenticationRepositoryInterface
	EmployeeProfileRepositoryInterface
	EmployeeEmploymentRepositoryInterface
	EmployeeProvisioningRepoInterface
	EmployeeRefreshSessionRepositoryInterface
}

type EmployeeMainRepositoryInterface interface {

	// this methods creates the core employee record with basic identity details (name, email, phone) and role.
	CreateEmployee(ctx context.Context, emp *coreEmployeeModels.Employee) error

	// this retuens the employee record of employee with given email. This is used for authentication and identity resolution.
	GetEmployeeByID(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) (*coreEmployeeModels.Employee, error)
	// Search employee by email across the org. Used for authentication and identity resolution.
	GetEmployeeByEmail(ctx context.Context, email string) (*coreEmployeeModels.Employee, error)
	EmployeeExistsByEmail(ctx context.Context, email string) (bool, error)
	ListEmployeesByOrg(ctx context.Context, orgID uuid.UUID) ([]*coreEmployeeModels.Employee, error)
	UpdateEmployee(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, updates map[string]any) error
}

type EmployeeAuthenticationRepositoryInterface interface {
	CreateEmployeeAuth(ctx context.Context, auth *coreEmployeeModels.EmployeeAuth) error
	GetAuthByEmployeeID(ctx context.Context, empID uuid.UUID) (*coreEmployeeModels.EmployeeAuth, error)
	GetByUsername(ctx context.Context, username string) (*coreEmployeeModels.EmployeeAuth, error)
	UpdateLastLogin(ctx context.Context, empID uuid.UUID, at time.Time) error
}

type EmployeeOnboardingRepositoryInterface interface {
	CreateEmployeeOnBoarding(ctx context.Context, onboarding *coreEmployeeModels.EmployeeOnboarding) error
	GetEmplOnBoardingByEmployeeID(ctx context.Context, empID uuid.UUID) (*coreEmployeeModels.EmployeeOnboarding, error)
	UpdateStatus(ctx context.Context, empID uuid.UUID, status string) error
}

type EmployeeProfileRepositoryInterface interface {
	CreateEmployeeProfile(ctx context.Context, profile *coreEmployeeModels.EmployeeProfile) error
	GetProfileByEmployeeID(ctx context.Context, empID uuid.UUID) (*coreEmployeeModels.EmployeeProfile, error)
	UpdateEmployeeProfile(ctx context.Context, profile *coreEmployeeModels.EmployeeProfile) error
}

type EmployeeEmploymentRepositoryInterface interface {
	CreateEmployement(ctx context.Context, emp *coreEmployeeModels.EmployeeEmployment) error
	GetByEmployeeID(ctx context.Context, empID uuid.UUID) (*coreEmployeeModels.EmployeeEmployment, error)
}

// EmployeeRefreshSessionRepositoryInterface backs mobile refresh-token
// rotation. RotateRefreshSession is the only method that may advance a
// session's jti, and it does so atomically — see its doc comment.
type EmployeeRefreshSessionRepositoryInterface interface {
	CreateRefreshSession(ctx context.Context, session *coreEmployeeModels.EmployeeRefreshSession) error
	GetRefreshSessionByID(ctx context.Context, sessionID uuid.UUID) (*coreEmployeeModels.EmployeeRefreshSession, error)
	RotateRefreshSession(ctx context.Context, sessionID, expectedJTI, newJTI uuid.UUID, now time.Time) (bool, error)
	RevokeRefreshSession(ctx context.Context, sessionID uuid.UUID, reason string, at time.Time) error
	RevokeAllRefreshSessionsForEmployee(ctx context.Context, employeeID uuid.UUID, reason string, at time.Time) (int64, error)
	PurgeExpiredRefreshSessions(ctx context.Context, before time.Time) (int64, error)
}

type EmployeeProvisioningRepoInterface interface {
	CreateEmployeeProvision(ctx context.Context, orgID uuid.UUID, provisions []coreEmployeeModels.EmployeeProvisioning) error
	GetEmployeeProvioionByEmployeeID(ctx context.Context, orgID uuid.UUID, employeeID uuid.UUID) ([]coreEmployeeModels.EmployeeProvisioning, error)
	AddConsumeTimeInProvisioningToken(ctx context.Context, orgID uuid.UUID, employeeID uuid.UUID, consumptionTime time.Time) error
}
