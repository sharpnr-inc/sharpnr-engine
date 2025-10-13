package signupservice

import (
	"context"
	"fmt"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/pkg/jwtpack"
	"time"
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
	existingOrg.Country = req.Address.Country // or you can pass it via DTO if dynamic
	existingOrg.ISRegistered = req.Registration.IsRegister

	if req.Registration.IsRegister {
		existingOrg.GstNumber = &req.Registration.GSTNumber
		existingOrg.CinNumber = &req.Registration.CINNumber
	} else {
		na := "N/A"
		existingOrg.GstNumber = &na
		existingOrg.CinNumber = &na
	}

	existingOrg.IsRegistrationComplete = true

	existingOrg.UpdatedAt = time.Now()

	if err := s.orgRepo.UpdateOrg(ctx, existingOrg); err != nil {
		return nil, "", fmt.Errorf("failed to update organization details: %w", err)
	}

	authToken, err := jwtpack.GenerateJwtToken(existingOrg.RootEmail, jwtpack.TokenTypeAuth, existingOrg.SharpnrID, string(existingOrg.OrgOprationType))
	if err != nil {
		return nil, "", err
	}

	return &dto.Org_Complete_Regesistration_RESPONSE_DTO{}, authToken, nil
}
