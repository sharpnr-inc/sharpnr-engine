package signupservice

import (
	"context"
	"fmt"
	"sharpnr-engine/core/constants"
	"sharpnr-engine/core/core_modules_events"
	core_events_names "sharpnr-engine/core/core_modules_events/coreEventsNames"
	dto "sharpnr-engine/core/dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"sharpnr-engine/modules/hospitality/hospitality_setup"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"
	"sharpnr-engine/pkg/jwtpack"
	"time"

	"github.com/google/uuid"
)

func (s *signupService) CompleteOrgRegistration(
	ctx context.Context,
	req dto.Org_Complete_Registration_REQ_DTO,
	orgEmail, orgSharpnrID string,
) (*dto.Org_Complete_Registration_RESPONSE_DTO, string, error) {

	// -------------------- Step 1: Find Organization --------------------
	existingOrg, found, err := s.orgRepo.FindOrgBySharpnrID(ctx, orgSharpnrID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to fetch organization: %w", err)
	}
	if !found {
		return nil, "", fmt.Errorf("organization not found with ID: %s", orgSharpnrID)
	}

	// -------------------- Step 2: Update Org Info --------------------
	existingOrg.Address = req.Address.AddressLine
	existingOrg.City = req.Address.City
	existingOrg.State = req.Address.State
	existingOrg.Pincode = req.Address.Pincode
	existingOrg.Country = req.Address.Country
	existingOrg.IsRegistrationComplete = true
	existingOrg.UpdatedAt = time.Now()

	// -------------------- Step 3: Determine Role --------------------
	role := constants.Super_Admin
	if req.Is_owner {
		role = constants.Org_Owners
	}

	defaultKitchen := &hospitalitymodels.Kitchen{
		ID:             uuid.New(),
		OrganizationID: existingOrg.ID,
		Organization:   *existingOrg,
	}

	if err := s.kitchenRepo.CreateKitchen(ctx, defaultKitchen); err != nil {
		return nil, "", fmt.Errorf("failed to create kitchen: %w", err)
	}

	domainSteps := hospitality_setup.HospitalityModuleSetupFactory(req.OrgSubModule)

	domainCtx := hospitality_setup.SetupContext{
		Org:            existingOrg,
		Kitchen:        defaultKitchen,
		BarRepo:        s.barRepo,
		HotelRepo:      s.hotelRepo,
		RestaurantRepo: s.restaurantRepo,
		MenuRepo:       s.menuRepo,
		UIRenderRepo:   s.uiRenderRepo,
	}

	for _, step := range domainSteps {
		if err := step.Execute(ctx, domainCtx); err != nil {
			return nil, "", fmt.Errorf(
				"hospitality setup failed at step [%s]: %w",
				step.Name(),
				err,
			)
		}
	}

	// -------------------- Step 6: Update Org in DB --------------------
	if err := s.orgRepo.UpdateOrg(ctx, existingOrg); err != nil {
		return nil, "", fmt.Errorf("failed to update organization details: %w", err)
	}

	// -------------------- Step 7: Create Default Employee --------------------
	defEmp := &sharphr_models.Employee{
		ID:              uuid.New(),
		FirstName:       "__Default__Name",
		LastName:        "__Default__Last__Name",
		Email:           orgEmail,
		OrganizationID:  existingOrg.ID,
		Organization:    *existingOrg,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Role:            role,
		IsEmailVerified: true,
	}

	createdEmp, err := s.hrEmpRepo.CreateEmployee(ctx, defEmp)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create default employee: %w", err)
	}

	// -------------------- Step 8: Create Org Policy --------------------
	if err := s.policyRepo.CreatePolicyOfOrg(ctx, existingOrg.ID.String(), existingOrg.SharpnrID, createdEmp.ID.String()); err != nil {
		fmt.Printf("⚠️ Failed to create policy for org %s: %v\n", existingOrg.ID, err)
	}

	// -------------------- Step 9: Generate Auth Token --------------------
	empIDStr := createdEmp.ID.String()
	authToken, err := jwtpack.GenerateJwtToken(
		existingOrg.RootEmail,
		jwtpack.TokenTypeAuth,
		existingOrg.SharpnrID,
		string(existingOrg.OrgOperationType),
		&empIDStr,
	)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate auth token: %w", err)
	}

	// -------------------- Step 10: Build and Return Response --------------------
	response := &dto.Org_Complete_Registration_RESPONSE_DTO{
		OrgID:        existingOrg.SharpnrID,
		OrgName:      existingOrg.Name,
		EmployeeID:   createdEmp.ID,
		EmployeeName: createdEmp.FirstName,
		EmployeeRole: string(createdEmp.Role),
	}

	s.sere.Publish(ctx, core_events_names.OrgCreated, core_modules_events.OrganizationCreated{
		OrganizationID:   existingOrg.ID,
		OrganizationName: existingOrg.Name,
		OrganizationType: string(existingOrg.OrgOperationType),
		Modules:          req.OrgSubModule,
	})

	fmt.Printf("✅ Organization registration completed for: %s (%s)\n", existingOrg.Name, existingOrg.SharpnrID)
	return response, authToken, nil
}
