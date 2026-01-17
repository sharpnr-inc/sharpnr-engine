package org_signup_services

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	"sharpnr-engine/core/organization/models"
	"sharpnr-engine/shared/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c *orgSignUpService) StartOrgOnBoarding(
	ctx context.Context,
	stagingSlug string,
	req org_signup_dtos.OrgOnBoardingRequest_dto,
) (*org_signup_dtos.OrgOnBoardingSuccessResponse_dto, error) {

	var result *org_signup_dtos.OrgOnBoardingSuccessResponse_dto

	err := c.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		// 🔑 Inject tx into context
		txCtx := dbtx.InjectTx(ctx, tx)

		stagingOrg, err := c.tempOrgSignupRepo.Get(txCtx, stagingSlug)
		if err != nil {
			return err
		}

		newSharpID, err := utils.GenerateUniqueSharpnrOrgID(
			txCtx,
			c.mainOrgSignUpRepo,
		)
		if err != nil {
			return err
		}

		org := &models.Organization{
			ID:        uuid.New(),
			Name:      stagingOrg.Name,
			SharpnrID: newSharpID,
			// OrgOperationType: constants.OrgType(req.OrgSector),
			Status:    "CREATED",
			RootEmail: stagingOrg.Email,
			CreatedAt: time.Now(),
		}

		if err := c.mainOrgSignUpRepo.CreateOrganization(txCtx, org); err != nil {
			return err
		}

		orgAddr := &models.OrganizationOffice{
			ID:             uuid.New(),
			OrganizationID: org.ID,
			Code:           "HQ",
			Name:           stagingOrg.Name,
			IsHQ:           true,
			IsActive:       true,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := c.mainOrgSignUpRepo.AddOrgOfficeAddress(txCtx, orgAddr); err != nil {
			return err
		}

		orgSubc := &models.OrganizationSubscription{
			ID:             uuid.New(),
			OrganizationID: org.ID,
			PlanCode:       "FREE",
			Status:         "Active",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := c.mainOrgSignUpRepo.CreateOrgSubscription(txCtx, orgSubc); err != nil {
			return err
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

		if err := c.mainOrgSignUpRepo.AddOrgVerificationDetails(
			txCtx,
			orgVerification,
		); err != nil {
			return err
		}

		// 🔥 SERE runs INSIDE SAME TX
		wf, err := c.sere.ExecuteTx(
			txCtx,
			"ORG_ONBOARD_WORKFLOW",
			map[string]any{
				"org_id":         org.ID,
				"email":          stagingOrg.Email,
				"sector":         req.OrgSector,
				"sector_modules": req.OrgSectorModules,
				"hashPass":       stagingOrg.Password,
			},
		)
		if err != nil {
			return err
		}

		employeeID := wf.Data["employee_id"].(uuid.UUID)

		result = &org_signup_dtos.OrgOnBoardingSuccessResponse_dto{
			OrgId:        org.ID,
			EmployeeID:   employeeID,
			EmployeeRole: "SUPER_ADMIN",
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
