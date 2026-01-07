package org_signup_services

import (
	"context"
	"sharpnr-engine/core/constants"
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

	// Todo := First We need to create the Employee and then
	// *Step 2 := We need to use the Hospitality Setup function to Create the Hospitality and its module
	// *Step 3 := We need to Create the Accounting for the Hospitality by sending or Publishing the Event

	res := org_signup_dtos.OrgOnBoardingSuccessResponse_dto{
		OrgId:        org.ID,
		EmployeeID:   uuid.Nil,
		EmployeeRole: "",
	}

	return &res, err
}
