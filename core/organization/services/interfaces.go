package organization_services

import (
	"context"
	organization_constants "sharpnr-engine/core/organization/constants"
	organization_contracts "sharpnr-engine/core/organization/contracts"

	"github.com/google/uuid"
)

type OrganizationServiceInterface interface {
	OrganizationSignUpServiceInterface
	OrganizationGeneralServiceInterface
}

// this services is user for Signup process,
type OrganizationSignUpServiceInterface interface {
	CreateNewOrganization(ctx context.Context, input organization_contracts.InputDataForOrgAndEmp) (output organization_contracts.InputForCreateDefaultEmployee, err error)
	CreateDraftOrgForSignup(ctx context.Context, input organization_contracts.DraftOrganizationInput) (uuid.UUID, string, error)
	DiscardSignupStaging(ctx context.Context, tempID uuid.UUID) error
	GetSignupSlug(ctx context.Context, tempId uuid.UUID) (string, error)
	VerifySlugCreateDefaultOrg(ctx context.Context, otp string, slugID uuid.UUID) (output organization_contracts.InputDataForOrgAndEmp, err error)
}

// these are general services for organization, like status update, etc.
type OrganizationGeneralServiceInterface interface {
	IsOrgExists(ctx context.Context, rootEmail string) (bool, error)
	GetOrgDetailsByID(ctx context.Context, orgID, emplID uuid.UUID) (*organization_contracts.OrganizationResult, error)
	GetStatus(ctx context.Context, orgID uuid.UUID) (organization_constants.OrgSetupStatusType, error)
	UpdateStatus(ctx context.Context, orgID uuid.UUID, status organization_constants.OrgSetupStatusType) error
	VerifyOrgByID(ctx context.Context, orgID uuid.UUID) (bool, error)
	UpdateOrganizationName(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, name string) error
}

var (
	_ OrganizationSignUpServiceInterface = (*organizationService)(nil)
	_ OrganizationServiceInterface       = (*organizationService)(nil)
)
