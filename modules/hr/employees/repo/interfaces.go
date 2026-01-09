package hr_employee_repo

import (
	"context"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"
	"time"

	"github.com/google/uuid"
)

type EmployeeIdentityRepositoryInterface interface {
	Create(ctx context.Context, emp *sharphr_models.Employee) error
	GetByID(ctx context.Context, empID uuid.UUID) (*sharphr_models.Employee, error)
	GetByEmail(ctx context.Context, email string) (*sharphr_models.Employee, error)
	ListByOrg(ctx context.Context, orgID uuid.UUID) ([]*sharphr_models.Employee, error)
}

type EmployeeAuthRepositoryInterface interface {
	Create(ctx context.Context, auth *sharphr_models.EmployeeAuth) error
	GetByEmployeeID(ctx context.Context, empID uuid.UUID) (*sharphr_models.EmployeeAuth, error)
	GetByUsername(ctx context.Context, username string) (*sharphr_models.EmployeeAuth, error)
	UpdateLastLogin(ctx context.Context, empID uuid.UUID, at time.Time) error
}

type EmployeeOnboardingRepositoryInterface interface {
	Create(ctx context.Context, onboarding *sharphr_models.EmployeeOnboarding) error
	GetByEmployeeID(ctx context.Context, empID uuid.UUID) (*sharphr_models.EmployeeOnboarding, error)
	UpdateStatus(ctx context.Context, empID uuid.UUID, status string) error
}

type EmployeeVerificationRepositoryInterface interface {
	Create(ctx context.Context, v *sharphr_models.EmployeeVerification) error
	GetByTokenHash(ctx context.Context, hash string) (*sharphr_models.EmployeeVerification, error)
	MarkVerified(ctx context.Context, id uuid.UUID, at time.Time) error
}

type EmployeeProfileRepositoryInterface interface {
	Create(ctx context.Context, profile *sharphr_models.EmployeeProfile) error
	GetByEmployeeID(ctx context.Context, empID uuid.UUID) (*sharphr_models.EmployeeProfile, error)
	Update(ctx context.Context, profile *sharphr_models.EmployeeProfile) error
}

type EmployeeEmploymentRepositoryInterface interface {
	Create(ctx context.Context, emp *sharphr_models.EmployeeEmployment) error
	GetByEmployeeID(ctx context.Context, empID uuid.UUID) (*sharphr_models.EmployeeEmployment, error)
}

type EmployeePolicyRepositoryInterface interface {
	Create(ctx context.Context, policy *sharphr_models.EmployeePolicy) error
	GetByID(ctx context.Context, id uuid.UUID) (*sharphr_models.EmployeePolicy, error)
	GetByCode(ctx context.Context, orgID uuid.UUID, code string) (*sharphr_models.EmployeePolicy, error)
	ListByOrganization(ctx context.Context, orgID uuid.UUID) ([]*sharphr_models.EmployeePolicy, error)
}

type EmployeePolicyRuleRepositoryInterface interface {
	Create(ctx context.Context, rule *sharphr_models.EmployeePolicyRule) error

	ListByPolicyID(
		ctx context.Context,
		policyID uuid.UUID,
	) ([]*sharphr_models.EmployeePolicyRule, error)

	DeleteByPolicyID(
		ctx context.Context,
		policyID uuid.UUID,
	) error
}

type EmployeePrivilegePolicyRepositoryInterface interface {
	Assign(
		ctx context.Context,
		assignment *sharphr_models.EmployeePrivilegePolicy,
	) error

	ListByEmployeeID(
		ctx context.Context,
		employeeID uuid.UUID,
	) ([]*sharphr_models.EmployeePrivilegePolicy, error)

	RevokeByEmployeeID(
		ctx context.Context,
		employeeID uuid.UUID,
	) error
}
