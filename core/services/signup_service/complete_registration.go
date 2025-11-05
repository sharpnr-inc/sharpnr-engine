package signupservice

import (
	"context"
	"fmt"
	"sharpnr-engine/core/constants"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"sharpnr-engine/pkg/jwtpack"
	"sort"
	"strings"
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

	// -------------------- Step 4: Create UI Renderer --------------------
	// if _, err := s.uiRenderRepo.CreateOrgUIRenderer(ctx, existingOrg.ID, req.OrgSubModule, existingOrg.OrgOperationType); err != nil {
	// 	fmt.Printf("⚠️ Failed to create UI renderer for org %s: %v\n", existingOrg.ID, err)
	// }

	// -------------------- Step 5: Setup Module(s) --------------------
	modules := req.OrgSubModule
	sort.Strings(modules) // Ensure consistent order for key
	key := strings.Join(modules, "_")
	fmt.Printf("🔑 Module combination key: %s\n", key)

	switch key {
	case "BAR":
		if err := s.SetUpBar(ctx, existingOrg); err != nil {
			return nil, "", fmt.Errorf("failed to set up BAR module: %w", err)
		}
	case "HOTEL":
		if err := s.SetUpHotel(ctx, existingOrg); err != nil {
			return nil, "", fmt.Errorf("failed to set up HOTEL module: %w", err)
		}
	case "RESTAURANT":
		if err := s.SetUpRestaurant(ctx, existingOrg); err != nil {
			return nil, "", fmt.Errorf("failed to set up RESTAURANT module: %w", err)
		}
	case "BAR_HOTEL":
		if err := s.SetUpHotelBar(ctx, existingOrg); err != nil {
			return nil, "", fmt.Errorf("failed to set up BAR_HOTEL modules: %w", err)
		}
	case "BAR_RESTAURANT":
		if err := s.SetUpBarRestaurant(ctx, existingOrg); err != nil {
			return nil, "", fmt.Errorf("failed to set up BAR_RESTAURANT modules: %w", err)
		}
	case "HOTEL_RESTAURANT":
		if err := s.SetUpHotelRestaurant(ctx, existingOrg); err != nil {
			return nil, "", fmt.Errorf("failed to set up HOTEL_RESTAURANT modules: %w", err)
		}
	case "BAR_HOTEL_RESTAURANT":
		if err := s.SetUpBarHotelRestaurant(ctx, existingOrg); err != nil {
			return nil, "", fmt.Errorf("failed to set up BAR_HOTEL_RESTAURANT modules: %w", err)
		}
	default:
		fmt.Printf("⚠️ Unknown module combination: %s\n", key)
	}

	defaultKitchen := &hospitalitymodels.Kitchen{
		ID:             uuid.New(),
		OrganizationID: existingOrg.ID,
		Organization:   *existingOrg,
	}
	if err := s.kitchenRepo.CreateKitchen(ctx, defaultKitchen); err != nil {
		return nil, "", fmt.Errorf("failed to create kitchen: %w", err)
	}

	// -------------------- Step 6: Update Org in DB --------------------
	if err := s.orgRepo.UpdateOrg(ctx, existingOrg); err != nil {
		return nil, "", fmt.Errorf("failed to update organization details: %w", err)
	}

	// -------------------- Step 7: Create Default Employee --------------------
	defEmp := &models.Employee{
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

	createdEmp, err := s.empRepo.CreateEmployee(ctx, defEmp)
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

	fmt.Printf("✅ Organization registration completed for: %s (%s)\n", existingOrg.Name, existingOrg.SharpnrID)
	return response, authToken, nil
}
