package org_signup_services

import (
	"context"
	"sharpnr-engine/core/constants"
	hospitality_org_models "sharpnr-engine/core/hospitality-org-factory/models"
	"sharpnr-engine/core/models"
	org_signup_dtos "sharpnr-engine/core/org-signup/dtos"
	"sharpnr-engine/shared/utils"
	"time"

	"github.com/google/uuid"
)

func (c *orgSignUpService) StartOrgOnBoarding(
	ctx context.Context,
	stagingSlug string,
	req org_signup_dtos.OrgOnBoardingRequest_dto,
) (*org_signup_dtos.OrgOnBoardingSuccessResponse_dto, error) {

	// Step One to get the Details to Verify

	stagingOrg, err := c.tempOrgSignupRepo.Get(ctx, stagingSlug)
	if err != nil {
		return nil, err
	}

	newSharpID, err := utils.GenerateUniqueSharpnrOrgID(ctx, c.mainOrgSignUpRepo)
	if err != nil {
		return nil, err
	}

	// then We Move to Find What Kind organization is this
	org := &models.Organization{
		ID:               uuid.New(),
		Name:             stagingOrg.Name,
		SharpnrID:        newSharpID,
		OrgOperationType: constants.OrgType(req.OrgSector),
		Status:           "CREATED",
		RootEmail:        stagingOrg.Email,
		CreatedAt:        time.Now(),
	}

	if err := c.mainOrgSignUpRepo.CreateOrganization(ctx, org); err != nil {
		return nil, err
	}

	orgAddr := &models.OrganizationOffice{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Code:           "HQ",
		Name:           stagingOrg.Name,
		AddressLine:    req.OrgAddressLine,
		City:           req.OrgCity,
		State:          req.OrgState,
		Pincode:        req.OrgPincode,
		Country:        req.OrgCountry,
		IsHQ:           true,
		IsActive:       true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := c.mainOrgSignUpRepo.AddOrgOfficeAddress(ctx, *orgAddr); err != nil {
		return nil, err
	}

	orgSubc := &models.OrganizationSubscription{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		PlanCode:       "FREE",
		Status:         "Active",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := c.mainOrgSignUpRepo.CreateOrgSubscription(ctx, *orgSubc); err != nil {
		return nil, err
	}

	orgVerification := &models.OrganizationVerification{
		ID:               uuid.New(),
		OrganizationID:   org.ID,
		EmailVerified:    true,
		MobileVerified:   false,
		PhysicalVerified: false,
		GSTVerified:      false,
		CINVerified:      false,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err := c.mainOrgSignUpRepo.AddOrgVerificationDetails(ctx, *orgVerification); err != nil {
		return nil, err
	}

	switch req.OrgSector {
	case "HOSPITALITY":
		hospOrg := hospitality_org_models.OrganizationHospitality{
			ID:             uuid.New(),
			OrganizationID: org.ID,
			Status:         "CREATED",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		if err := c.hospitalityFactoryRepo.CreateHospitalityOrg(ctx, &hospOrg); err != nil {
			return nil, err
		}

	}

	// c.hrEmployeeService.CreateDefaultSuperAdmin(ctx, org.ID)

	res := org_signup_dtos.OrgOnBoardingSuccessResponse_dto{
		OrgId:        org.ID,
		EmployeeID:   uuid.Nil,
		EmployeeRole: "",
	}

	return &res, err
}
