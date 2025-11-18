package restaurantservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

func (r *restaurantService) UpdateRestaurantGeneralDetails(ctx context.Context, req hospitalitydtos.Restaurant_UpdateOrEdit_General_Details_REQ, orgId uuid.UUID) (bool, error) {
	// Step 1: Fetch existing restaurant by orgId
	restaurant, err := r.restaurantRepo.GetRestaurant(ctx, orgId)
	if err != nil {
		return false, err
	}

	// Step 2: Map DTO → Model
	restaurant.LegalName = req.HotelLegalName
	restaurant.DisplayName = req.HotelDisplayName

	restaurant.FSSAINumber = req.FSSAINumber
	restaurant.GSTNumber = req.GSTNumber
	restaurant.CINNumber = req.CINNumber

	restaurant.Address = req.AddressLine
	restaurant.City = req.City
	restaurant.State = req.State
	restaurant.Pincode = req.Pincode
	restaurant.Country = req.Country // Add this field to model if needed

	restaurant.ContactEmail = req.EmailAddress
	restaurant.ContactNumber = req.ContactNumber
	restaurant.Manager_name = req.ManagerName

	// Step 3: Save changes
	err = r.restaurantRepo.UpdateRestaurant(ctx, restaurant)
	if err != nil {
		return false, err
	}

	return true, nil
}
