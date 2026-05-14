package organization_services

import (
	"context"
	organization_constants "sharpnr-engine/core/organization/constants"
	organization_contracts "sharpnr-engine/core/organization/contracts"

	"github.com/google/uuid"
)

type OrganizationServiceInterface interface {
	CreateNewOrganization(ctx context.Context, input organization_contracts.InputDataForOrgAndEmp) (output organization_contracts.InputForCreateDefaultEmployee, err error)
	CreateDraftOrgForSignup(ctx context.Context, input organization_contracts.DraftOrganizationInput) (uuid.UUID, string, error)

	IsOrgExists(ctx context.Context, rootEmail string) (bool, error)

	VerifyOrgByID(ctx context.Context, orgID uuid.UUID) (bool, error)

	GetStatus(ctx context.Context, orgID uuid.UUID) (organization_constants.OrgSetupStatusType, error)

	UpdateStatus(ctx context.Context, orgID uuid.UUID, status organization_constants.OrgSetupStatusType) error
	UpdateOrganizationName(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, name string) error

	GetSignupSlug(ctx context.Context, tempId uuid.UUID) (string, error)

	VerifySlugCreateDefaultOrg(ctx context.Context, otp string, slugID uuid.UUID) (output organization_contracts.InputDataForOrgAndEmp, err error)

	AddHQ(ctx context.Context, input organization_contracts.OrganizationHQ) error
	GetOrgDetailsByID(ctx context.Context, orgID, emplID uuid.UUID) (*organization_contracts.OrganizationResult, error)
}
