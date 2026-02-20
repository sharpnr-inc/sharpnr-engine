package kitchenservices

import (
	"context"
	"errors"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	kitchen_contracts "sharpnr-engine/modules/hospitality/kitchen/contracts"
	restaurant_events "sharpnr-engine/modules/hospitality/restaurant/events"
	sse_engine "sharpnr-engine/sse"

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

	// 3. Broadcast updated list to all subscribers of this kitchen
	// We run this in a goroutine to keep the API response fast
	go func() {
		// Using Background context as the request context might be cancelled
		bgCtx := context.Background()

		// Fetch the full updated list of KOTs for this kitchen
		updatedKots, err := s.GetKots(bgCtx, ticket.KitchenID)
		if err != nil {
			return // Log: "Broadcast fetch failed"
		}

		// Push to the same engine used by the controllers
		s.sseEngine.MessageChan <- sse_engine.Event{
			Topic:   fmt.Sprintf("kitchen_%s", ticket.KitchenID.String()),
			Payload: updatedKots,
		}
	}()

	return nil

}
