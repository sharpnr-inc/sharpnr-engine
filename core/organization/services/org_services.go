package organization_services

import (
	"context"
	"errors"
	"fmt"
	organization_constants "sharpnr-engine/core/organization/constants"
	organization_contracts "sharpnr-engine/core/organization/contracts"
	"sharpnr-engine/core/organization/models"
	"sharpnr-engine/shared/utils"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This return tell org Exist or not
func (r *organizationService) IsOrgExists(ctx context.Context, rootEmail string) (bool, error) {
	// 1. Attempt to fetch the organization by email
	_, err := r.orgMainRepo.GetOrganizationByEmail(ctx, rootEmail)

	if err != nil {
		// 2. If the error is "Record Not Found", the org does NOT exist (Return false, nil)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		// 3. If it's a real database error (connection, etc.), return the error
		return false, fmt.Errorf("failed to check organization existence: %w", err)
	}

	// 4. If err is nil, it means a record was found
	return true, nil
}

/* ---------------------------------- <***> --------------------------------- */

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
	sharpID, err := utils.GenerateUniqueSharpnrOrgID(ctx, r.orgMainRepo)
	if err != nil {
		return output, err
	}

	// 2. Build organization model
	org := &models.Organization{
		ID:        orgID,
		Name:      "Default Organization",
		SharpnrID: sharpID,
		Status:    string(organization_constants.OrgNotConfigured),
		RootEmail: input.Email,
	}

	// 3. Persist organization
	if err := r.orgMainRepo.CreateOrganization(ctx, org); err != nil {
		return output, err
	}

	orgStatus := &models.OrganizationSetupStatus{
		ID:              uuid.New(),
		OrganizationID:  org.ID,
		IsSetupComplete: false,
	}

	if err := r.orgMainRepo.CreateOrganizationStatus(ctx, orgStatus); err != nil {
		return output, err
	}

	// 4. Prepare data for employee creation
	output = organization_contracts.InputForCreateDefaultEmployee{
		OrgID:      orgID,
		Email:      org.RootEmail,
		HashedPass: input.HashedPass,
	}

	return output, nil
}
