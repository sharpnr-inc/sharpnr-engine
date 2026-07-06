package kitchenservices

import (
	"context"
	"errors"
	kitchen_contracts "sharpnr-engine/business_ops/hospitality/kitchen/contracts"
	restaurant_events "sharpnr-engine/business_ops/hospitality/restaurant/events"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

		// 2. Explicitly query items for THIS ticket
		items, err := s.kitchenRepo.GetKitchenTicketItemsByTicketID(ctx, t.ID)
		if err != nil {
			continue
		}

		for _, itm := range items {
			itemDTO := kitchen_contracts.KOTItem_resp_dto{
				OrderID:      itm.ID,
				MenuItemID:   itm.MenuItemID,
				MenuItemName: itm.MenuItemName,
				Quantity:     itm.Quantity,
				Remarks:      itm.Remarks,
				IsPrinted:    itm.IsPrinted,
				Addons:       []kitchen_contracts.KOTAddon_resp_dto{},
			}

			// --- VARIANT LOGIC ---
			// Check if this item has a variant assigned
			variant, err := s.kitchenRepo.GetKitchenTicketItemVariant(ctx, itm.ID)
			if err == nil {
				// Record found: Populate pointers
				itemDTO.VaraintName = &variant.VariantName
				itemDTO.VariantID = &variant.VaraintID
			} else if !errors.Is(err, gorm.ErrRecordNotFound) {
				// A real database error occurred
				return nil, err
			}
			// If ErrRecordNotFound, we do nothing; itemDTO fields stay nil

			// --- ADDONS LOGIC ---
			addons, err := s.kitchenRepo.GetKitchenTicketItemAddOn(ctx, itm.ID)
			if err == nil {
				for _, a := range addons {
					itemDTO.Addons = append(itemDTO.Addons, kitchen_contracts.KOTAddon_resp_dto{
						ID:       a.ID,
						AddOnID:  a.AddOnID,
						Quantity: a.Quantity,
					})
				}
			} else if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}

			kotDTO.Items = append(kotDTO.Items, itemDTO)
		}

		resp = append(resp, kotDTO)
	}

	return resp, nil
}

func (s *kitchenService) UpdateKotStatus(ctx context.Context, kitchenID uuid.UUID, kotID uuid.UUID, status string) error {

	// 1. Get the ticket first to know which KitchenID it belongs to
	// (Required for the SSE Topic)
	ticket, err := s.kitchenRepo.GetKOTByID(ctx, kotID)
	if err != nil {
		return err
	}

	err = s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)
		if err := s.kitchenRepo.UpdateKotStatus(txCtx, kotID, status); err != nil {
			return err
		}

		payload := restaurant_events.OrderStatusUpdatePayload{
			OrderID: ticket.OrderID,
			Status:  status,
		}

		if err := s.seroEngine.Publish(txCtx, "ORDER_STATUS_UPDATED", payload); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	// 3. Broadcast updated list to all subscribers of this kitchen without
	// delaying the API response.
	go s.broadcastKots(ticket.KitchenID)

	return nil

}
