package organization_services

import (
	"context"
	organization_contracts "sharpnr-engine/core/organization/contracts"

	"github.com/google/uuid"
)

type OrganizationServiceInterface interface {

	// This method gets the Organization By ID
	// GetOrgByID(
	// 	ctx context.Context,
	// 	orgId uuid.UUID,
	// ) (organization_contracts.OrganizationResult, error)

	// //
	// GetOrgBySharpID(
	// 	ctx context.Context,
	// 	sharpID string,
	// ) (organization_contracts.OrganizationResult, error)

	//
	IsOrgExists(
		ctx context.Context,
		rootEmail string,
		rootMobile string,
	) (bool, error)

	GetOrgHQ()

	// This is Creates a base organization
	CreateOrganization(
		ctx context.Context,
		input interface{},
	) error

	AddHQ(
		ctx context.Context,
		input interface{},
	) error

	CreateDraftOrgForSignup(
		ctx context.Context,
		input organization_contracts.DraftOrganizationInput,
	) (string, string, error)

	GetSignupSlug(
		ctx context.Context, tempId uuid.UUID,
	) (string, error)

	VerifySlugCreateDefaultOrg(
		ctx context.Context,
		otp string,
		slugID uuid.UUID,
	) (output organization_contracts.InputDataForOrgAndEmp, err error)

	CreateNewOrganization(
		ctx context.Context,
		input organization_contracts.InputDataForOrgAndEmp,
	) (output organization_contracts.InputForCreateDefaultEmployee, err error)
}
