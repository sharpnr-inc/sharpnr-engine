package signupservice

import (
	"context"
	"errors"
	"fmt"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"
	"sharpnr-engine/shared/mailer"
	"sharpnr-engine/shared/utils"
)

func (s *signupService) InitiateOrgRegistration(ctx context.Context, req dto.Org_Signup_REQUEST_DTO) (*dto.Org_Signup_RESPONSE_DTO, string, error) {
	// Step 1 we will check first that if data is there
	if req.Name == "" || req.Email == "" || req.Mobile == "" || req.Password == "" {
		return nil, "", errors.New("all fields are required")
	}

	orgEmail := req.Email
	orgPhone := req.Mobile

	// Step 1 → Check if email or phone exists in Redis
	existsInRedis, err := s.orgRepo.EmailOrPhoneExistsInRedis(ctx, orgEmail, orgPhone)
	if err != nil {
		return nil, "", fmt.Errorf("failed to check redis: %w", err)
	}

	if existsInRedis {
		// Update OTP in redis and send a NEW OTP mail
		fmt.Println("ORG Email Exits")
		ok, otp, err := mailer.VerifyOtp_Org_Registration(orgEmail)
		if !ok || err != nil {
			return nil, "", fmt.Errorf("failed to send otp email: %w", err)
		}

		token, err := utils.GenerateOTPToken(orgEmail)
		if err != nil {
			return nil, "", err
		}

		// Save otp token in redis
		update := &models.TempOrgRegistrationData{
			Email: orgEmail,
			Phone: orgPhone,
			OTP:   otp,
		}
		if err := s.orgRepo.CreateOrgInRedisDB(ctx, update); err != nil {
			return nil, "", fmt.Errorf("failed to update redis otp: %w", err)
		}

		return &dto.Org_Signup_RESPONSE_DTO{
			Token: token,
		}, "", nil
	}

	// Step 2 → Not in Redis → Check MainDB
	existsInMainDB, err := s.orgRepo.EmailOrPhoneExistsInMainDB(ctx, orgEmail, orgPhone)
	if err != nil {
		return nil, "", fmt.Errorf("failed to check database: %w", err)
	}
	if existsInMainDB {
		return nil, "", fmt.Errorf("email or phone already registered in main database")
	}

	// Step 3 → Create temp org in Redis and send OTP mail
	ok, otpToken, err := mailer.VerifyOtp_Org_Registration(orgEmail)
	if !ok || err != nil {
		return nil, "", fmt.Errorf("failed to send otp email: %w", err)
	}

	tempOrg := &models.TempOrgRegistrationData{
		Name:  req.Name,
		Email: orgEmail,
		Phone: orgPhone,
		// OrgType:  req.OrgType,
		Password: req.Password,
		OTP:      otpToken,
	}

	token, err := utils.GenerateOTPToken(orgEmail)
	if err != nil {
		return nil, "", err
	}

	if err := s.orgRepo.CreateOrgInRedisDB(ctx, tempOrg); err != nil {
		return nil, "", fmt.Errorf("failed to create temp org in redis: %w", err)
	}

	return &dto.Org_Signup_RESPONSE_DTO{
		Token: token,
	}, "", nil
}
