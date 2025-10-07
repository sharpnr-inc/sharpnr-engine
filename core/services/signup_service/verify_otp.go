package signupservice

import (
	"context"
	"fmt"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"
	"sharpnr-engine/pkg/jwtpack"

	// policyengine "sharpnr-engine/policy_engine"
	"sharpnr-engine/shared/utils"

	"github.com/google/uuid"
)

// Verify OTP Service
func (s *signupService) VerifyOrgRegistration(ctx context.Context, email, code string) (*dto.Org_Verify_Otp_RESPONSE_DTO, bool, string, error) {
	// 1. Parse JWT claims

	fmt.Println("Finding Data for the %w", email)
	tempOrg, err := s.orgRepo.GetOrgFromRedisDB(ctx, email)
	if err != nil {
		return nil, false, "", fmt.Errorf("failef to fetch org data: %w", err)
	}

	if tempOrg == nil {
		return nil, false, "", fmt.Errorf("no pending registration found org: %w", err)
	}

	if tempOrg.OTP != code {
		return nil, false, "", fmt.Errorf("invalid OTP code :%w", err)
	}

	if err := s.orgRepo.DeleteOrgFromRedisDB(ctx, email); err != nil {
		return nil, false, "", fmt.Errorf("failed to cleanup redis: %w", err)
	}

	newOrgId, err := utils.GenerateUniqueSharpnrOrgID(ctx, s.orgRepo)
	if err != nil {
		return nil, false, "", fmt.Errorf("failed to generate unique org id: %w", err)
	}

	org := &models.Organization{
		ID:                     uuid.New(),
		SharpnrID:              newOrgId,
		Name:                   tempOrg.Name,
		RootEmail:              tempOrg.Email,
		MasterPhoneNumber:      tempOrg.Phone,
		Password:               tempOrg.Password,
		OrgOprationType:        tempOrg.OrgType,
		IsRegistrationComplete: false,
	}

	if err := s.orgRepo.CreateOrgInMainDB(ctx, org); err != nil {
		return nil, false, "", fmt.Errorf("failed to save organization in main DB: %w", err)
	}

	newToken, err := jwtpack.GenerateJwtToken(org.RootEmail, jwtpack.TokenTypeVerifyOTP, org.SharpnrID, string(org.OrgOprationType))
	if err != nil {
		return nil, false, "", fmt.Errorf("failed to create Auth Token %w", err)
	}

	// 5. Build response DTO
	res := &dto.Org_Verify_Otp_RESPONSE_DTO{
		Name:       org.Name,
		Sharpnr_Id: org.SharpnrID,
		Email:      org.RootEmail,
		ORG_TYPE:   org.OrgOprationType,
	}

	return res, true, newToken, nil
}
