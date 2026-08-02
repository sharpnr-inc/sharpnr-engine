package organization_services

import (
	"context"
	"crypto/subtle"
	"errors"
	organization_constants "sharpnr-engine/core/organization/constants"
	organization_contracts "sharpnr-engine/core/organization/contracts"
	"sharpnr-engine/core/organization/models"
	"sharpnr-engine/shared/utils"
	"strings"

	"github.com/google/uuid"
)

func (r *organizationService) GetSignupSlug(
	ctx context.Context, tempID uuid.UUID,
) (string, error) {

	// 2. Fetch from Redis
	staging, err := r.orgRedisRepo.Get(ctx, tempID)
	if err != nil {
		return "", err
	}

	// Same two failure modes as VerifySlugCreateDefaultOrg, so they carry the
	// same sentinels: the caller has to tell "your signup expired, start over"
	// apart from "Redis is down", and a bare errors.New cannot express that.
	if staging == nil {
		return "", ErrSignupNotFound
	}

	// 3. Return ONLY email
	if staging.Email == "" {
		return "", ErrInvalidSignupData
	}

	return staging.Email, nil
}

/* ---------------------------------- <***> --------------------------------- */

// DiscardSignupStaging removes a consumed signup staging record (and its
// email/phone reverse-lookup keys) from Redis.
//
// This is separate from VerifySlugCreateDefaultOrg on purpose: Redis takes no
// part in the database transaction, so the record may only be dropped once that
// transaction has actually committed. Call it after the commit, never inside.
func (r *organizationService) DiscardSignupStaging(
	ctx context.Context, tempID uuid.UUID,
) error {
	return r.orgRedisRepo.Delete(ctx, tempID)
}

/* ---------------------------------- <***> --------------------------------- */

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

	staging, err := r.orgRedisRepo.Get(ctx, slugID)
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

/* ---------------------------------- <***> --------------------------------- */

func (r *organizationService) CreateDraftOrgForSignup(
	ctx context.Context,
	input organization_contracts.DraftOrganizationInput,
) (signupSlug uuid.UUID, otp string, err error) {

	if input.Email == "" && (input.Mobile == nil || *input.Mobile == "") {
		return uuid.Nil, "", errors.New("email or phone is required")
	}

	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return uuid.Nil, "", err
	}

	stagingID := uuid.New()

	otp, err = utils.GenerateOTP()
	if err != nil {
		return uuid.Nil, "", err
	}

	var phone string
	if input.Mobile != nil {
		phone = *input.Mobile
	}

	tempOrg := &models.OrganizationSignUpStaging{
		ID:       stagingID,
		Email:    strings.ToLower(strings.TrimSpace(input.Email)),
		Phone:    phone, // safe
		Otp:      otp,
		Password: hashPassword,
	}

	if err := r.orgRedisRepo.Create(ctx, tempOrg); err != nil {
		return uuid.Nil, "", err
	}

	return stagingID, otp, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *organizationService) CreateNewOrganization(
	ctx context.Context,
	input organization_contracts.InputDataForOrgAndEmp,
) (output organization_contracts.InputForCreateDefaultEmployee, err error) {

	// 1. Generate OrgID (SOURCE OF TRUTH)
	orgID := uuid.New()
	sharpID, err := utils.GenerateUniqueSharpnrOrgID(ctx, r.orgMainRepo)
	if err != nil {
		return output, err
	}

	// 2. Build organization model
	org := &models.Organization{
		ID:        orgID,
		Name:      "Default Organization",
		SharpnrID: sharpID,
		Status:    organization_constants.OrgConfigStatusConfigNotStarted,
		RootEmail: input.Email,
	}

	// 3. Persist organization
	if err := r.orgMainRepo.CreateOrganization(ctx, org); err != nil {
		return output, err
	}

	// orgStatus := &models.OrganizationSetupStatus{
	// 	ID:              uuid.New(),
	// 	OrganizationID:  org.ID,
	// 	IsSetupComplete: false,
	// }

	// if err := r.orgMainRepo.CreateOrganizationStatus(ctx, orgStatus); err != nil {
	// 	return output, err
	// }

	// 4. Prepare data for employee creation
	output = organization_contracts.InputForCreateDefaultEmployee{
		OrgID:      orgID,
		Email:      org.RootEmail,
		HashedPass: input.HashedPass,
	}

	return output, nil
}

/* ---------------------------------- <***> --------------------------------- */
