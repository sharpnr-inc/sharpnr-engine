package menu_sere_adapter

import (
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"
	"sharpnr-engine/sere"
	"time"

	"github.com/google/uuid"
)

func barDefaultMenuSetup(
	menuRepo menurepo.MenuRepositoryInterface,
) sere.WorkflowHandler {

	return func(wf *sere.Workflow) error {

		// ------------------------------------------------
		// 0️⃣ Check if BAR is selected
		// ------------------------------------------------
		modules := wf.Data["sector_modules"].([]string)

		hasBar := false
		for _, m := range modules {
			if m == "BAR" {
				hasBar = true
				break
			}
		}

		if !hasBar {
			return nil // 🚫 BAR not requested
		}

		// ------------------------------------------------
		// 1️⃣ Preconditions
		// ------------------------------------------------
		orgID := wf.Data["org_id"].(uuid.UUID)

		// ------------------------------------------------
		// 2️⃣ Create BAR menu
		// ------------------------------------------------
		barMenu := &hospitalityMenu_models.MenuGroup{
			ID:             uuid.New(),
			OrganizationID: orgID,
			Name:           "Bar Menu",
			Type:           "BAR",
			IsActive:       true,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := menuRepo.CreateMenuGroup(wf.Ctx, barMenu); err != nil {
			return err
		}

		return nil
	}
}

func hotelDefaultMenuSetup(menuRepo menurepo.MenuRepositoryInterface) sere.WorkflowHandler {
	return func(wf *sere.Workflow) error {

		// ------------------------------------------------
		// 0️⃣ Check if HOTEL is selected in request
		// ------------------------------------------------

		modules := wf.Data["sector_modules"].([]string)

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
		orgID := wf.Data["org_id"].(uuid.UUID)

		hotelMenu := &hospitalityMenu_models.MenuGroup{
			ID:             uuid.New(),
			OrganizationID: orgID,
			Name:           "Hotel Menu",
			Type:           "HOTEL",
			IsActive:       true,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := menuRepo.CreateMenuGroup(wf.Ctx, hotelMenu); err != nil {
			return err
		}

		return nil
	}
}

func restaurantDefaultMenuSetup(
	menuRepo menurepo.MenuRepositoryInterface,
) sere.WorkflowHandler {

	return func(wf *sere.Workflow) error {

		// ------------------------------------------------
		// 0️⃣ Check if RESTAURANT is selected
		// ------------------------------------------------
		modules := wf.Data["sector_modules"].([]string)

		hasRestaurant := false
		for _, m := range modules {
			if m == "RESTAURANT" {
				hasRestaurant = true
				break
			}
		}

		if !hasRestaurant {
			return nil // 🚫 RESTAURANT not requested
		}

		// ------------------------------------------------
		// 1️⃣ Preconditions
		// ------------------------------------------------
		orgID := wf.Data["org_id"].(uuid.UUID)

		// ------------------------------------------------
		// 2️⃣ Create RESTAURANT menu
		// ------------------------------------------------
		restaurantMenu := &hospitalityMenu_models.MenuGroup{
			ID:             uuid.New(),
			OrganizationID: orgID,
			Name:           "Restaurant Menu",
			Type:           "RESTAURANT",
			IsActive:       true,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := menuRepo.CreateMenuGroup(wf.Ctx, restaurantMenu); err != nil {
			return err
		}

		return nil
	}
}
