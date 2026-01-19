package organization_services

import (
	"context"
	"errors"
	organization_contracts "sharpnr-engine/core/organization/contracts"

	"github.com/google/uuid"
)

// VerifySlugCreateDefaultOrg implements OrganizationServiceInterface.
// !! Remind to change the name of this method to VerifySlugAndOtp
func (r *organizationService) VerifySlugCreateDefaultOrg(
	ctx context.Context,
	otp string,
	slugID uuid.UUID,
) (output organization_contracts.InputDataForOrgAndEmp, err error) {

	slug := slugID.String()

	staging, err := r.orgRedisRepo.Get(ctx, slug)
	if err != nil {
		return output, err
	}

	if staging == nil {
		return output, errors.New("signup not found or expired")
	}

	if staging.Otp != otp {
		return output, errors.New("invalid otp")
	}

	if staging.Email == "" || staging.Password == "" {
		return output, errors.New("invalid signup data")
	}

	output = organization_contracts.InputDataForOrgAndEmp{
		Email:      staging.Email,
		HashedPass: staging.Password,
		Phone:      staging.Phone,
	}

	if err := r.orgRedisRepo.Delete(ctx, slug); err != nil {
		return output, err
	}

	return output, nil
}
