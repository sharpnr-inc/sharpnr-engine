package kitchenservices

import (
	"context"
	kitchen_contracts "sharpnr-engine/modules/hospitality/kitchen/contracts"

	"github.com/google/uuid"
)

func (s *kitchenService) GetKots(ctx context.Context, kitchenID uuid.UUID) ([]kitchen_contracts.KOT_resp_dto, error) {
	// 1. Fetch the Ticket Headers first
	tickets, err := s.kitchenRepo.GetKitchenTicketsByKitchenID(ctx, kitchenID)
	if err != nil {
		return nil, err
	}

	var resp []kitchen_contracts.KOT_resp_dto

	for _, t := range tickets {
		kotDTO := kitchen_contracts.KOT_resp_dto{
			ID:        t.ID,
			TicketNo:  t.TicketNo,
			Status:    t.Status,
			CreatedAt: t.CreatedAt,
			Items:     []kitchen_contracts.KOTItem_resp_dto{},
		}

		// 2. Explicitly query items for THIS ticket (Flat DB style)
		items, err := s.kitchenRepo.GetKitchenTicketItemsByTicketID(ctx, t.ID)
		if err != nil {
			continue // Or handle error
		}

		for _, itm := range items {
			itemDTO := kitchen_contracts.KOTItem_resp_dto{
				ID:         itm.ID,
				MenuItemID: itm.MenuItemID,
				VariantID:  itm.VariantID,
				Quantity:   itm.Quantity,
				Remarks:    itm.Remarks,
				IsPrinted:  itm.IsPrinted,
				Addons:     []kitchen_contracts.KOTAddon_resp_dto{},
			}

			// 3. Explicitly query addons for THIS item
			addons, err := s.kitchenRepo.GetKitchenTicketItemAddOn(ctx, itm.ID)
			if err == nil {
				for _, a := range addons {
					itemDTO.Addons = append(itemDTO.Addons, kitchen_contracts.KOTAddon_resp_dto{
						ID:       a.ID,
						AddOnID:  a.AddOnID,
						Quantity: a.Quantity,
					})
				}
			}

			kotDTO.Items = append(kotDTO.Items, itemDTO)
		}

		resp = append(resp, kotDTO)
	}

	return resp, nil
}
