package restaurantservices

import (
	"context"
	"fmt"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

// Helper function to update string fields safely
func updateStringIfPresent(target *string, source *string) {
	if source != nil {
		*target = *source
	}
}

func (r *restaurantService) UpdateRestaurantGeneralDetails(
	ctx context.Context,
	req hospitalitydtos.Restaurant_UpdateOrEdit_General_Details_REQ,
	orgId uuid.UUID,
) (bool, error) {
	// Step 1: Validate restaurant_id
	restaurantUUID, err := uuid.Parse(req.RestaurantId)
	if err != nil {
		return false, fmt.Errorf("invalid restaurant_id: %w", err)
	}

	// Step 2: Fetch existing restaurant by orgId
	restaurant, err := r.restaurantRepo.GetRestaurantByOrgId(ctx, orgId)
	if err != nil {
		return false, fmt.Errorf("failed to fetch restaurant: %w", err)
	}

	// Step 3: Verify restaurant belongs to organization
	if restaurant.ID != restaurantUUID {
		return false, fmt.Errorf("restaurant does not belong to organization")
	}

	// Step 4: Update only provided fields (nil fields are skipped)
	updateStringIfPresent(&restaurant.LegalName, req.RestaurantLegalName)
	updateStringIfPresent(&restaurant.DisplayName, req.RestaurantDisplayName)
	updateStringIfPresent(&restaurant.FSSAINumber, req.FSSAINumber)
	updateStringIfPresent(&restaurant.GSTNumber, req.GSTNumber)
	updateStringIfPresent(&restaurant.CINNumber, req.CINNumber)
	updateStringIfPresent(&restaurant.Address, req.AddressLine)
	updateStringIfPresent(&restaurant.City, req.City)
	updateStringIfPresent(&restaurant.State, req.State)
	updateStringIfPresent(&restaurant.Country, req.Country)
	updateStringIfPresent(&restaurant.Pincode, req.Pincode)
	updateStringIfPresent(&restaurant.ContactEmail, req.EmailAddress)
	updateStringIfPresent(&restaurant.ContactNumber, req.ContactNumber)
	updateStringIfPresent(&restaurant.ManagerName, req.ManagerName)

	// Step 5: Save changes to database
	err = r.restaurantRepo.UpdateRestaurant(ctx, restaurant)
	if err != nil {
		return false, fmt.Errorf("failed to update restaurant: %w", err)
	}

	return true, nil
}
