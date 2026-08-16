package organization_repo

import (
	"context"
	organization_constants "sharpnr-engine/core/organization/constants"
	"sharpnr-engine/core/organization/models"

	"github.com/google/uuid"
)

type RedisOrganizationRepositoryInterface interface {

	// Store temp org with TTL
	Create(ctx context.Context, tempOrg *models.OrganizationSignUpStaging) error
	// Fetch temp org
	Get(ctx context.Context, signupID uuid.UUID) (*models.OrganizationSignUpStaging, error)
	// Remove after success or expiry
	Delete(ctx context.Context, signupID uuid.UUID) error
	GetByEmailOrPhone(ctx context.Context, email string, phone string) (bool, error)
}

type OrganizationRepositoryInterface interface {
	CreateOrganization(ctx context.Context, org *models.Organization) error
	GetOrganizationByID(ctx context.Context, orgID uuid.UUID) (*models.Organization, error)
	GetOrganizationBySharpID(ctx context.Context, sharpID string) (*models.Organization, error)
	GetOrganizationByEmail(ctx context.Context, rootEmail string) (*models.Organization, error)
	UpdateOrganizationStatus(ctx context.Context, orgID uuid.UUID, status organization_constants.OrgSetupStatusType) error
	UpdateOrgName(ctx context.Context, orgID uuid.UUID, name string) error
	CreateOrgSubscription(ctx context.Context, orgSub *models.OrganizationSubscription) error
	GetOrgSubscriptionPlan(ctx context.Context, orgId uuid.UUID) (*models.OrganizationSubscription, error)
	AddOrgVerificationDetails(ctx context.Context, orgVerification *models.OrganizationVerification) error
}
