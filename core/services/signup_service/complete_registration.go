package signupservice

import (
	"context"
	"fmt"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"
	"sharpnr-engine/pkg/jwtpack"
	"time"

	"github.com/google/uuid"
)

func (s *signupService) CompleteOrgRegistration(
	ctx context.Context,
	req dto.Org_Complete_Regesistration_REQ_DTO,
	org_email, org_sharpnr_id string,
) (*dto.Org_Complete_Regesistration_RESPONSE_DTO, string, error) {

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
	existingOrg.Country = "India" // or you can pass it via DTO if dynamic

	existingOrg.GstNumber = &req.Registration.GSTNumber
	existingOrg.CinNumber = &req.Registration.CINNumber
	existingOrg.ISRegistered = req.Registration.IsRegister
	existingOrg.IsRegistrationComplete = true

	existingOrg.UpdatedAt = time.Now()

	if err := s.orgRepo.UpdateOrg(ctx, existingOrg); err != nil {
		return nil, "", fmt.Errorf("failed to update organization details: %w", err)
	}

	bankAccounts := make([]models.OrganizationBankAccount, len(req.BankDetails))
	for i, b := range req.BankDetails {
		bankAccounts[i] = models.OrganizationBankAccount{
			ID:             uuid.New(),
			OrganizationID: existingOrg.ID, // assign orgID here
			BankName:       b.BankName,
			AccountNumber:  b.AccountNumber,
			IFSCCode:       b.IFSCCode,
			BranchName:     b.BranchName,
			IsPrimary:      b.IsPrimary,
			IsSecondary:    b.IsSecondary,
		}
	}

	err = s.orgRepo.CreateBankAccounts(ctx, existingOrg.ID, bankAccounts)
	if err != nil {
		return nil, "", err
	}

	DefaultSuperAdmin := &models.Employee{
		ID:              uuid.New(),
		FirstName:       req.SuperAdmin.Name,
		Email:           existingOrg.RootEmail,
		OrganizationID:  existingOrg.ID,
		CreatedAt:       time.Now(),
		Role:            "SUPER_ADMIN",
		IsEmailVerified: true,
	}

	_, fault := s.empRepo.CreateEmployee(ctx, DefaultSuperAdmin)
	if fault != nil {
		return nil, "", fmt.Errorf("failed to create super admin: %w", fault)
	}

	// policy, err := policyengine.CreateDefaultPolicy(existingOrg)

	authToken, err := jwtpack.GenerateJwtToken(existingOrg.RootEmail, jwtpack.TokenTypeAuth, existingOrg.SharpnrID, string(existingOrg.OrgOprationType))
	if err != nil {
		return nil, "", err
	}

	return &dto.Org_Complete_Regesistration_RESPONSE_DTO{}, authToken, nil
}
