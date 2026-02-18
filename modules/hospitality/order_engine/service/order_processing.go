package hospitality_order_engine_service

import (
	"context"
	kitchen_contracts "sharpnr-engine/modules/hospitality/kitchen/contracts"
	restaurant_events "sharpnr-engine/modules/hospitality/restaurant/events"

	"github.com/google/uuid"
)

func (s *hospitalityOrderEngineService) ProcessOrder(ctx context.Context, data restaurant_events.OrderCreatedPayload) error {

	// 1. Group items by KitchenID
	// Key: KitchenID (uuid.UUID), Value: Slice of KOT items for that specific kitchen
	groups := make(map[uuid.UUID][]kitchen_contracts.KOTItem_dto)

	for _, itm := range data.Items {
		// Get the KitchenID from the Menu Service
		kitchenID, err := s.menuService.GetKitchenOfMenuItem(ctx, itm.MenuItemID)
		if err != nil {
			// If an item isn't assigned to a kitchen, skip it to prevent transaction failure.
			// You might want to route this to a 'Default' kitchen ID here.
			continue
		}

		// Map Variants from the Event DTO to the Kitchen DTO
		var kotVariants []kitchen_contracts.KOTExtra_dto
		for _, v := range itm.Variants { // Changed from SelectedVariants
			kotVariants = append(kotVariants, kitchen_contracts.KOTExtra_dto{
				ID:   v.VariantID,
				Name: v.VariantName,
				Qty:  1, 
			})
		}

		// Map Add-ons from the Event DTO to the Kitchen DTO
		var kotAddons []kitchen_contracts.KOTExtra_dto
		for _, a := range itm.Addons { // Changed from SelectedAddOns
			kotAddons = append(kotAddons, kitchen_contracts.KOTExtra_dto{
				ID:   a.AddOnID,
				Name: a.Name,
				Qty:  a.Qty, // Using Qty from the event DTO
			})
		}

		// Prepare the KOT Item DTO
		kotItem := kitchen_contracts.KOTItem_dto{
			OrderItemID: itm.ID,         // The PK of the RestaurantOrderItem
			MenuItemID:  itm.MenuItemID,
			Quantity:    itm.Quantity,
			Remarks:     itm.Notes,
			Variants:    kotVariants,
			Addons:      kotAddons,
		}

		groups[kitchenID] = append(groups[kitchenID], kotItem)
	}

	// 2. Dispatch to Kitchen Service for each group
	for kID, kotItems := range groups {
		if kID == uuid.Nil {
			continue
		}

		req := kitchen_contracts.CreateKOT_req_dto{
			OrganizationID: data.OrganizationID,
			PropertyID:     data.RestaurantID,
			KitchenID:      kID,
			OrderID:        data.OrderID,
			Items:          kotItems,
		}

		// Hand off to Kitchen Service to create the tickets
		if err := s.kitchenService.CreateKOT(ctx, req); err != nil {
			return err
		}
	}

	return nil
}
