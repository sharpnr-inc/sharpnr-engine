package kitchenservices

import (
	"context"
	kitchen_contracts "sharpnr-engine/modules/hospitality/kitchen/contracts"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"
	"time"

	"github.com/google/uuid"
)

func (s *kitchenService) CreateKOT(ctx context.Context, input kitchen_contracts.CreateKOT_req_dto) error {
	now := time.Now()

	// 1. Create the Ticket Header (KOT)
	ticket := &kitchen_models.KitchenTicket{
		ID:             uuid.New(),
		OrganizationID: input.OrganizationID,
		PropertyID:     input.PropertyID,
		KitchenID:      input.KitchenID,
		Status:         "PENDING",
		TicketNo:       "KOT-AUTO-1", // Replace with your sequence logic
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := s.kitchenRepo.CreateKitchenTicket(ctx, ticket); err != nil {
		return err
	}

	// 2. Iterate through line items
	for _, itm := range input.Items {

		// Note: Using itm.Variants[0].ID if your model only supports one variant.
		// If multiple variants are needed, the model KitchenTicketItem needs updating.
		var variantID *uuid.UUID
		if len(itm.Variants) > 0 {
			vID := itm.Variants[0].ID
			variantID = &vID
		}

		ticketItem := &kitchen_models.KitchenTicketItem{
			ID:              uuid.New(),
			KitchenTicketID: ticket.ID,
			MenuItemID:      itm.MenuItemID,
			VariantID:       variantID,
			Quantity:        int(itm.Quantity),
			Remarks:         itm.Remarks,
			IsPrinted:       false,
			CreatedAt:       now,
			UpdatedAt:       now,
		}

		if err := s.kitchenRepo.CreateKitchenTicketItem(ctx, ticketItem); err != nil {
			return err
		}

		// 3. Handle Addons for this specific item
		for _, adn := range itm.Addons {
			addonRecord := &kitchen_models.KitchenTicketItemAddon{
				ID:                  uuid.New(),
				KitchenTicketItemID: ticketItem.ID,
				AddOnID:             adn.ID,
				Quantity:            adn.Qty,
				CreatedAt:           now,
			}

			if err := s.kitchenRepo.CreateKitchenTicketItemAddon(ctx, addonRecord); err != nil {
				return err
			}
		}
	}

	return nil
}
