package hotel_sere_adapter

import (
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"
	"sharpnr-engine/sere"

	"github.com/google/uuid"
)

func defaultHotelHandler(
	hotelRepo hotelrepo.HotelRepositoryInterface,
) sere.WorkflowHandler {
	return func(w *sere.Workflow) error {

		// ------------------------------------------------
		// 0️⃣ Check if HOTEL is selected in request
		// ------------------------------------------------
		modules := w.Data["sector_modules"].([]string)

		hasHotel := false
		for _, m := range modules {
			if m == "HOTEL" {
				hasHotel = true
				break
			}
		}

		if !hasHotel {
			return nil // 🚫 HOTEL not requested
		}

		// ------------------------------------------------
		// 1️⃣ Preconditions
		// ------------------------------------------------

		orgId := w.Data["org_id"].(uuid.UUID)
		kitchenId := w.Data["kitchen_id"].(uuid.UUID)

		// ------------------------------------------------
		// 2️⃣ Idempotency via registry table
		// ------------------------------------------------

		defaultHotel := &hotel_models.Hotel{
			ID:             uuid.New(),
			OrganizationID: orgId,
			Name:           "system Created Default",
			IsKitchen:      true,
			KitchenID:      &kitchenId,
		}

		if err := hotelRepo.CreateHotel(w.Ctx, defaultHotel); err != nil {
			return err
		}

		return nil
	}
}
