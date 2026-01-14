package restaurant_sere_adapters

import (
	hospitality_org_factory_repository "sharpnr-engine/core/hospitality-org-factory/repository"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
	"sharpnr-engine/sere"
	"time"

	"github.com/google/uuid"
)

func createDefaultRestaurant(
	restaurantRepo restaurantrepo.RestaurantRepository,
	hospitalityOrgFactoryRepository hospitality_org_factory_repository.HospitalityOrgFactoryRepoInterface,
) sere.WorkflowHandler {
	return func(wf *sere.Workflow) error {

		// ------------------------------------------------
		// 0️⃣ Check if HOTEL is selected in request
		// ------------------------------------------------
		modules := wf.Data["sector_modules"].([]string)

		hasRestaurant := false
		for _, m := range modules {
			if m == "HOTEL" {
				hasRestaurant = true
				break
			}
		}

		if !hasRestaurant {
			return nil // 🚫 HOTEL not requested
		}

		// ------------------------------------------------
		// 1️⃣ Preconditions
		// ------------------------------------------------
		orgID := wf.Data["org_id"].(uuid.UUID)
		kitchenID := wf.Data["kitchen_id"].(uuid.UUID)

		defaultRes := &restaurant_models.Restaurant{
			ID:             uuid.New(),
			Name:           "Default Restaurant",
			OrganizationID: orgID,
			IsKitchen:      true,
			KitchenID:      (*uuid.UUID)(&kitchenID),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := restaurantRepo.CreateRestaurant(wf.Ctx, defaultRes); err != nil {
			return err
		}

		return nil
	}
}
