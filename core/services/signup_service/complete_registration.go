package signupservice

import (
	"context"
	"fmt"
	"sharpnr-engine/core/constants"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"
	"sharpnr-engine/pkg/jwtpack"

	// sharpnr_ui_renderer "sharpnr-engine/sharpnr_ui_renderer"
	"time"

	"github.com/google/uuid"
)

func (s *signupService) CompleteOrgRegistration(
	ctx context.Context,
	req dto.Org_Complete_Registration_REQ_DTO,
	org_email, org_sharpnr_id string,
) (*dto.Org_Complete_Registration_RESPONSE_DTO, string, error) {

	existingOrg, found, err := s.orgRepo.FindOrgBySharpnrID(ctx, org_sharpnr_id)
	if err != nil {
		return nil, "", err
	}
	if !found {
		return nil, "", fmt.Errorf("organization not found with ID: %s", org_sharpnr_id)
	}
	existingOrg.Address = req.Address.AddressLine
	existingOrg.City = req.Address.City
	existingOrg.State = req.Address.State
	existingOrg.Pincode = req.Address.Pincode
	existingOrg.Country = req.Address.Country
	existingOrg.IsRegistrationComplete = true
	existingOrg.UpdatedAt = time.Now()

	role := constants.Super_Admin
	if req.Is_owner {
		role = constants.Org_Owners
	}

	defemp := &models.Employee{
		ID:              uuid.New(),
		FirstName:       "__Default__Name",
		LastName:        "__Default__Last__Name",
		Email:           org_email,
		OrganizationID:  existingOrg.ID,
		Organization:    *existingOrg,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Role:            role,
		IsEmailVerified: true,
	}

	empIdStr := defemp.ID.String()

	if err := s.orgRepo.UpdateOrg(ctx, existingOrg); err != nil {
		return nil, "", fmt.Errorf("failed to update organization details: %w", err)
	}

	DefaultSuperAdmin, err := s.empRepo.CreateEmployee(ctx, defemp)
	if err != nil {
		return nil, "", nil
	}

	if err := s.policyRepo.CreatePolicyOfOrg(ctx, existingOrg.ID.String(), existingOrg.SharpnrID, DefaultSuperAdmin.ID.String()); err != nil {
		fmt.Println("Failed to create policy:", err)
	}

	authToken, err := jwtpack.GenerateJwtToken(existingOrg.RootEmail, jwtpack.TokenTypeAuth, existingOrg.SharpnrID, string(existingOrg.OrgOperationType), &empIdStr)
	if err != nil {
		return nil, "", err
	}

	return &dto.Org_Complete_Registration_RESPONSE_DTO{
		OrgID:        existingOrg.SharpnrID,
		OrgName:      existingOrg.Name,
		EmployeeID:   DefaultSuperAdmin.ID,
		EmployeeName: DefaultSuperAdmin.FirstName,
		EmployeeRole: string(DefaultSuperAdmin.Role),
	}, authToken, nil
}
