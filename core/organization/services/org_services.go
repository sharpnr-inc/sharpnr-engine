package organization_services

import (
	"context"
	"errors"
	organization_contracts "sharpnr-engine/core/organization/contracts"
	"sharpnr-engine/core/organization/models"
	"sharpnr-engine/shared/utils"
	"strings"

	"github.com/google/uuid"
)

// This return tell org Exist or not
func (r *organizationService) IsOrgExists(
	ctx context.Context,
	rootEmail string,
	mobile string,
) (bool, error) {

	emailExist, err := r.orgMainRepo.ExistsByEmail(ctx, rootEmail)
	if err != nil {
		return false, err
	}

	if emailExist {
		return true, nil
	}

	// phoneExist, err := r.orgMainRepo.ExistsByPhone(ctx, mobile)
	// if err != nil {
	// 	return false, err
	// }

	// if phoneExist {
	// 	return true, nil
	// }

	return false, nil
}

func (r *organizationService) CreateDraftOrgForSignup(
	ctx context.Context,
	input organization_contracts.DraftOrganizationInput,
) (signupSlug string, otp string, err error) {

	if input.Email == "" && (input.Mobile == nil || *input.Mobile == "") {
		return "", "", errors.New("email or phone is required")
	}

	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return "", "", err
	}

	stagingID := uuid.New()

	otp, err = utils.GenerateOTP()
	if err != nil {
		return "", "", err
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
		return "", "", err
	}

	return stagingID.String(), otp, nil
}

//----------------------------------------------------------------------------------------------------------------//

func (r *organizationService) CreateNewOrganization(
	ctx context.Context,
	input organization_contracts.InputDataForOrgAndEmp,
) (output organization_contracts.InputForCreateDefaultEmployee, err error) {

	// 1. Generate OrgID (SOURCE OF TRUTH)
	orgID := uuid.New()

	// 2. Build organization model
	org := &models.Organization{
		ID:        orgID,
		Name:      "Default Organization", // or derive from email later
		SharpnrID: orgID.String(),         // or your own generator
		Status:    "CREATED",
		RootEmail: input.Email,
	}

	// 3. Persist organization
	if err := r.orgMainRepo.CreateOrganization(ctx, org); err != nil {
		return output, err
	}

	// 4. Prepare data for employee creation
	output = organization_contracts.InputForCreateDefaultEmployee{
		OrgID:      orgID,
		Email:      input.Email,
		HashedPass: input.HashedPass,
	}

	return output, nil
}
