package signupservice

import (
	"context"
	"fmt"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"
	"sharpnr-engine/shared/utils"

	"github.com/google/uuid"
)

// Verify OTP Service
func (s *signupService) VerifyOrgRegistration(ctx context.Context, token string, code string) (*dto.Org_Verify_Otp_RESPONSE_DTO, bool, error) {
	// 1. Parse JWT claims
	claims, err := utils.ParseJWTToken(token)
	if err != nil {
		return nil, false, fmt.Errorf("invalid or expired token: %w", err)
	}

	if claims["tokenType"] != nil && claims["tokenType"] != "OTP_VERIFICATION" {
		return nil, false, fmt.Errorf("invalid token type: expected otp")
	}

	// 2. Extract email
	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return nil, false, fmt.Errorf("email not found in token")
	}

	fmt.Println("Finding Data for the %w", email)
	tempOrg, err := s.orgRepo.GetOrgFromRedisDB(ctx, email)
	if err != nil {
		return nil, false, fmt.Errorf("failef to fetch org data: %w", err)
	}

	if tempOrg == nil {
		return nil, false, fmt.Errorf("no pending registration found org: %w", err)
	}

	if tempOrg.OTP != code {
		return nil, false, fmt.Errorf("invalid OTP code :%w", err)
	}

	if err := s.orgRepo.DeleteOrgFromRedisDB(ctx, email); err != nil {
		return nil, false, fmt.Errorf("failed to cleanup redis: %w", err)
	}

	newOrgId, err := utils.GenerateUniqueSharpnrOrgID(ctx, s.orgRepo)
	if err != nil {
		return nil, false, fmt.Errorf("failed to generate unique org id: %w", err)
	}

	org := &models.Organization{
		ID:                uuid.New(),
		SharpnrID:         newOrgId,
		Name:              tempOrg.Name,
		RootEmail:         tempOrg.Email,
		MasterPhoneNumber: tempOrg.Phone,
		Password:          tempOrg.Password,
		OrgOprationType:   tempOrg.OrgType,
	}

	if err := s.orgRepo.CreateOrgInMainDB(ctx, org); err != nil {
		return nil, false, fmt.Errorf("failed to save organization in main DB: %w", err)
	}

	newToken, err := utils.GenerateAuthToken(org.RootEmail)
	if err != nil {
		return nil, false, fmt.Errorf("failed to generate auth token: %w", err)
	}

	// 5. Build response DTO
	res := &dto.Org_Verify_Otp_RESPONSE_DTO{
		Name:       org.Name,
		Sharpnr_Id: org.SharpnrID,
		Email:      org.RootEmail,
		ORG_TYPE:   org.OrgOprationType,
		Token:      newToken,
	}

	return res, true, nil
}
