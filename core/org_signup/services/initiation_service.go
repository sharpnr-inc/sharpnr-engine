package org_signup_services

import (
	"context"
	"errors"
	"sharpnr-engine/core/models"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	"sharpnr-engine/shared/utils"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"github.com/google/uuid"
)

var (
	ErrSignupAlreadyInitiated = errors.New("signup already initiated")
	ErrOrgEmailExists         = errors.New("organization email already exists")
	ErrOrgPhoneExists         = errors.New("organization phone already exists")
	ErrInvalidOTP             = errors.New("invalid or expired otp")
)

// InitiateSignup implements OrgSignUpServiceInterface.
func (o *orgSignUpService) InitiateSignup(
	ctx context.Context,
	req org_signup_dtos.InitiateOrgSignUpRequest_Dto,
) (string, error) {

	// 1. Check staging (OTP already sent?)
	exists, err := o.tempOrgSignupRepo.GetByEmailOrPhone(
		ctx,
		req.OrgEmail,
		req.OrgMobile,
	)
	if err != nil {
		return "", err
	}
	if exists {
		return "", ErrSignupAlreadyInitiated
	}

	// 2. Check main DB (email)
	if ok, err := o.mainOrgSignUpRepo.ExistsByEmail(ctx, req.OrgEmail); err != nil {
		return "", err
	} else if ok {
		return "", ErrOrgEmailExists
	}

	// // 3. Check main DB (phone)
	// if ok, err := o.mainOrgSignUpRepo.ExistsByPhone(ctx, req.OrgMobile); err != nil {
	// 	return "", err
	// } else if ok {
	// 	return "", ErrOrgPhoneExists
	// }

	// 4. Generate OTP
	otp, err := utils.GenerateOTP()
	if err != nil {
		return "", err
	}

	// 5. Hash password
	hashedPassword, err := utils.HashPassword(req.OrgPassword)
	if err != nil {
		return "", err
	}

	// 6. Create staging record
	stageOrg := &models.OrganizationSignUpStaging{
		ID:       uuid.New(),
		Name:     req.OrgName,
		Email:    req.OrgEmail,
		Phone:    req.OrgMobile,
		Otp:      otp,
		Password: hashedPassword,
	}

	if err := o.tempOrgSignupRepo.Create(ctx, stageOrg); err != nil {
		return "", err
	}

	otpMailData := map[string]interface{}{
		"Code":   stageOrg.Otp,
		"Expiry": "10 minutes",
	}

	o.mailerEngine.Queues[sharp_mailer_engine.OTP] <- sharp_mailer_engine.EmailJob{
		To:       stageOrg.Email,
		Subject:  "Your Shapnr Verification Code",
		Template: sharp_mailer_engine.TmplVerification,
		Data:     otpMailData,
	}
	// 8. Return signup session ID (for cookie)
	return stageOrg.ID.String(), nil
}
