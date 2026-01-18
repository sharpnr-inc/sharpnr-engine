package employee_repo

import (
	"context"
	employee_models "sharpnr-engine/modules/hr/employees/models"
	"time"

	"github.com/google/uuid"
)

type EmployeeIdentityRepositoryInterface interface {
	Create(
		ctx context.Context,
		emp *employee_models.Employee,
	) error

	GetByID(
		ctx context.Context,
		empID uuid.UUID,
	) (*employee_models.Employee, error)

	GetByEmail(
		ctx context.Context,
		email string,
	) (*employee_models.Employee, error)

	ExistsByEmail(
		ctx context.Context,
		email string,
	) (bool, error)

	ListByOrg(
		ctx context.Context,
		orgID uuid.UUID,
	) ([]*employee_models.Employee, error)
}

type EmployeeAuthRepositoryInterface interface {
	Create(ctx context.Context, auth *employee_models.EmployeeAuth) error
	GetByEmployeeID(ctx context.Context, empID uuid.UUID) (*employee_models.EmployeeAuth, error)
	GetByUsername(ctx context.Context, username string) (*employee_models.EmployeeAuth, error)
	UpdateLastLogin(ctx context.Context, empID uuid.UUID, at time.Time) error
}

type EmployeeOnboardingRepositoryInterface interface {
	Create(ctx context.Context, onboarding *employee_models.EmployeeOnboarding) error
	GetByEmployeeID(ctx context.Context, empID uuid.UUID) (*employee_models.EmployeeOnboarding, error)
	UpdateStatus(ctx context.Context, empID uuid.UUID, status string) error
}

type EmployeeVerificationRepositoryInterface interface {
	Create(ctx context.Context, v *employee_models.EmployeeVerification) error
	GetByTokenHash(ctx context.Context, hash string) (*employee_models.EmployeeVerification, error)
	MarkVerified(ctx context.Context, id uuid.UUID, at time.Time) error
}

type EmployeeProfileRepositoryInterface interface {
	Create(ctx context.Context, profile *employee_models.EmployeeProfile) error
	GetByEmployeeID(ctx context.Context, empID uuid.UUID) (*employee_models.EmployeeProfile, error)
	Update(ctx context.Context, profile *employee_models.EmployeeProfile) error
}

type EmployeeEmploymentRepositoryInterface interface {
	Create(ctx context.Context, emp *employee_models.EmployeeEmployment) error
	GetByEmployeeID(ctx context.Context, empID uuid.UUID) (*employee_models.EmployeeEmployment, error)
}

type EmployeePolicyRepositoryInterface interface {
	Create(ctx context.Context, policy *employee_models.EmployeePolicy) error
	GetByID(ctx context.Context, id uuid.UUID) (*employee_models.EmployeePolicy, error)
	GetByCode(ctx context.Context, orgID uuid.UUID, code string) (*employee_models.EmployeePolicy, error)
	ListByOrganization(ctx context.Context, orgID uuid.UUID) ([]*employee_models.EmployeePolicy, error)
}

type EmployeePolicyRuleRepositoryInterface interface {
	Create(ctx context.Context, rule *employee_models.EmployeePolicyRule) error

	ListByPolicyID(
		ctx context.Context,
		policyID uuid.UUID,
	) ([]*employee_models.EmployeePolicyRule, error)

	DeleteByPolicyID(
		ctx context.Context,
		policyID uuid.UUID,
	) error
}

type EmployeePrivilegePolicyRepositoryInterface interface {
	Assign(
		ctx context.Context,
		assignment *employee_models.EmployeePrivilegePolicy,
	) error

	ListByEmployeeID(
		ctx context.Context,
		employeeID uuid.UUID,
	) ([]*employee_models.EmployeePrivilegePolicy, error)

	RevokeByEmployeeID(
		ctx context.Context,
		employeeID uuid.UUID,
	) error
}
