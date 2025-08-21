package services

import (
	"context"
	"fmt"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/interfaces"
	"sharpnr-engine/core/models"
	"sharpnr-engine/shared/mailer"
	"sharpnr-engine/shared/utils"
)

type orgRegistrationService struct {
	repo interfaces.IOrgRegistrationRepository
}

func OrgRegService(repo interfaces.IOrgRegistrationRepository) interfaces.OrgRegistrationService {
	return &orgRegistrationService{repo: repo}
}

func (s *orgRegistrationService) InitiateOrgRegistration(ctx context.Context, req dto.OrgInitialRegistration_Request_Dto) (*dto.OrgInitialRegistration_Response_Dto, string, error) {

	orgEmail := req.Email
	orgPhone := req.Phone

	// Step 1 → Check if email or phone exists in Redis
	existsInRedis, err := s.repo.EmailOrPhoneExistsInRedis(ctx, orgEmail, orgPhone)
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
		if err := s.repo.CreateOrgInRedisDB(ctx, update); err != nil {
			return nil, "", fmt.Errorf("failed to update redis otp: %w", err)
		}

		return &dto.OrgInitialRegistration_Response_Dto{
			Token: token,
		}, "", nil
	}

	// Step 2 → Not in Redis → Check MainDB
	existsInMainDB, err := s.repo.EmailOrPhoneExistsInMainDB(ctx, orgEmail, orgPhone)
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
		Name:     req.Name,
		Email:    orgEmail,
		Phone:    orgPhone,
		OrgType:  req.OrgType,
		Password: req.Password,
		OTP:      otpToken,
	}

	token, err := utils.GenerateOTPToken(orgEmail)
	if err != nil {
		return nil, "", err
	}

	if err := s.repo.CreateOrgInRedisDB(ctx, tempOrg); err != nil {
		return nil, "", fmt.Errorf("failed to create temp org in redis: %w", err)
	}

	return &dto.OrgInitialRegistration_Response_Dto{
		Token: token,
	}, "", nil
}
