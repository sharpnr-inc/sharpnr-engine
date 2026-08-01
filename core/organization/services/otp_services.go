package organization_services

import (
	"context"
	"crypto/subtle"
	"errors"
	organization_contracts "sharpnr-engine/core/organization/contracts"

	"github.com/google/uuid"
)

// Sentinel errors for the signup verification step. Callers match on these with
// errors.Is to decide what the client is told: these three are the caller's
// fault, anything else coming out of this function is an infrastructure failure
// and must not be reported as a rejected OTP.
var (
	ErrSignupNotFound    = errors.New("signup not found or expired")
	ErrInvalidOTP        = errors.New("invalid otp")
	ErrInvalidSignupData = errors.New("invalid signup data")
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
		return output, ErrSignupNotFound
	}

	if subtle.ConstantTimeCompare([]byte(staging.Otp), []byte(otp)) != 1 {
		return output, ErrInvalidOTP
	}

	if staging.Email == "" || staging.Password == "" {
		return output, ErrInvalidSignupData
	}

	output = organization_contracts.InputDataForOrgAndEmp{
		Email:      staging.Email,
		HashedPass: staging.Password,
		Phone:      staging.Phone,
	}

	// The staging record is deliberately NOT deleted here. This runs inside the
	// signup transaction, and Redis does not roll back with it: if a later step
	// (org, employee, sentinel, wizard) fails, the database unwinds while the
	// delete stands, leaving the user with a dead OTP and no way to retry.
	// The caller discards the record once the transaction has committed —
	// see DiscardSignupStaging.
	return output, nil
}
