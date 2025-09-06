package signupservice

import (
	"context"
	"fmt"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/shared/utils"
)

func (s *signupService) VerifyOrgRegistration(ctx context.Context, token string, code string) (*dto.Org_Verify_Otp_RESPONSE_DTO, bool, error) {
	// 1. Parse JWT claims
	claims, err := utils.ParseJWTToken(token)
	if err != nil {
		return nil, false, fmt.Errorf("invalid or expired token: %w", err)
	}

	// 2. Extract email
	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return nil, false, fmt.Errorf("email not found in token")
	}

	// 3. Extract OTP type (optional, only if you add it when generating)
	if claims["tokenType"] != nil && claims["tokenType"] != "otp" {
		return nil, false, fmt.Errorf("invalid token type: expected otp")
	}

	// 4. Validate OTP (here you’ll fetch from Redis instead of hardcoding)
	expectedCode := "123456" // placeholder
	if code != expectedCode {
		return nil, false, fmt.Errorf("invalid OTP code")
	}

	// 5. Build response DTO
	res := &dto.Org_Verify_Otp_RESPONSE_DTO{
		Email: email,
	}

	return res, true, nil
}
