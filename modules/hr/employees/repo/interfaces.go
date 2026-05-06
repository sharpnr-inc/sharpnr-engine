package employee_repo

import (
	"context"
	employee_models "sharpnr-engine/modules/hr/employees/models"
	"time"

	"github.com/google/uuid"
)

type EmployeeRepositoryInterface interface {
	EmployeeMainRepositoryInterface
	EmployeeVerificationRepositoryInterface
	EmployeeOnboardingRepositoryInterface
	EmployeeAuthenticationRepositoryInterface
	EmployeeVerificationRepositoryInterface
	EmployeeProfileRepositoryInterface
	EmployeeEmploymentRepositoryInterface
}

type EmployeeMainRepositoryInterface interface {
	CreateIdentity(
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

	UpdateIdentity(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, updates map[string]interface{}) error
}

type EmployeeAuthenticationRepositoryInterface interface {
	CreateEmployeeAuth(ctx context.Context, auth *employee_models.EmployeeAuth) error
	GetAuthByEmployeeID(ctx context.Context, empID uuid.UUID) (*employee_models.EmployeeAuth, error)
	GetByUsername(ctx context.Context, username string) (*employee_models.EmployeeAuth, error)
	UpdateLastLogin(ctx context.Context, empID uuid.UUID, at time.Time) error
}

type EmployeeOnboardingRepositoryInterface interface {
	CreateEmployeeOnBoarding(ctx context.Context, onboarding *employee_models.EmployeeOnboarding) error
	GetEmplOnBoardingByEmployeeID(ctx context.Context, empID uuid.UUID) (*employee_models.EmployeeOnboarding, error)
	UpdateStatus(ctx context.Context, empID uuid.UUID, status string) error
}

type EmployeeVerificationRepositoryInterface interface {
	CreateEmplVerification(ctx context.Context, v *employee_models.EmployeeVerification) error
	GetByTokenHash(ctx context.Context, hash string) (*employee_models.EmployeeVerification, error)
	MarkVerified(ctx context.Context, id uuid.UUID, at time.Time) error
}

type EmployeeProfileRepositoryInterface interface {
	CreateEmployeeProfile(ctx context.Context, profile *employee_models.EmployeeProfile) error
	GetProfileByEmployeeID(ctx context.Context, empID uuid.UUID) (*employee_models.EmployeeProfile, error)
	UpdateEmployeeProfile(ctx context.Context, profile *employee_models.EmployeeProfile) error

	CreateEmployeeSystemAccess(ctx context.Context, access *employee_models.EmployeeSystemAccess) (*employee_models.EmployeeSystemAccess, error)
	GetEmployeeSystemAccessByEmployeeID(ctx context.Context, empID uuid.UUID) ([]*employee_models.EmployeeSystemAccess, error)
	IsEmployeeHaveOnlineAccess(ctx context.Context, empID uuid.UUID) (bool, error)
}

type EmployeeEmploymentRepositoryInterface interface {
	CreateEmployement(ctx context.Context, emp *employee_models.EmployeeEmployment) error
	GetByEmployeeID(ctx context.Context, empID uuid.UUID) (*employee_models.EmployeeEmployment, error)
}
