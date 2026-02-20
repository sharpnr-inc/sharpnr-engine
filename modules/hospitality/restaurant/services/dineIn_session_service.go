package restaurantservices

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
	restaurant_events "sharpnr-engine/modules/hospitality/restaurant/events"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *restaurantService) StartDineSession(ctx context.Context,
	restaurantID uuid.UUID,
	orgID uuid.UUID,
	employeeID uuid.UUID,
	tableID uuid.UUID,
) (uuid.UUID, error) { // Removed the named 'diningSessionID' for cleaner return

	var sessionID uuid.UUID

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.restaurantRepo.UpdateTableStatus(txCtx, tableID, restaurant_constants.TableStatusOccupied); err != nil {
			return err
		}

		// 1. Generate the ID first so we can capture it
		sessionID = uuid.New()

		session := &restaurant_models.RestaurantDiningSession{
			ID:           sessionID,
			RestaurantID: restaurantID,
			TableID:      tableID,
			Status:       restaurant_constants.DineInSessionOpened,
			StartedBy:    employeeID,
			StartedAt:    time.Now(),
			CreatedAt:    time.Now(),
		}

		if err := s.restaurantRepo.CreateRestaurantDiningSession(txCtx, session); err != nil {
			return err
		}

		return nil
	})

	// 2. Check if the transaction failed
	if err != nil {
		return uuid.Nil, err
	}

	// 3. Return the captured ID
	return sessionID, nil
}

/* ----------------------------------- <*> ---------------------------------- */

func (s *restaurantService) AddOrdersInSession(
	ctx context.Context,
	dineInSessionID uuid.UUID,
	req *restaurant_contracts.DiningSessionOrder_req_dto,
	orgID uuid.UUID,
	restaurantID uuid.UUID,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. Initialize the Event Payload
		eventPayload := restaurant_events.OrderCreatedPayload{
			OrderID:        uuid.New(),
			OrganizationID: orgID,
			RestaurantID:   restaurantID,
			SessionID:      dineInSessionID,
			Items:          make([]restaurant_events.OrderItemEventDTO, 0),
		}

		// 2. Create Order Header
		ord := &restaurant_models.RestaurantOrders{
			ID:                        eventPayload.OrderID,
			OrganizationID:            orgID,
			/* ---------------- RestaurantID:              restaurantID, ---------------- */
			RestaurantDiningSessionID: &dineInSessionID,
			OrderNumber:               "1", // TODO Needed to fixed
			Status:                    "PENDING",
			CreatedAt:                 time.Now(),
			UpdatedAt:                 time.Now(),
		}

		if err := s.restaurantRepo.CreateOrder(txCtx, ord); err != nil {
			return err
		}

		for _, itm := range req.Items {

			// Build the Item Record for the DB
			itemRecord := &restaurant_models.RestaurantOrderItems{
				ID:                uuid.New(),
				RestaurantOrderID: ord.ID,
				MenuItemID:        itm.MenuItemID,
				MenuItemName:      itm.ItemName,
				Quantity:          itm.Quantity,
				UnitPrice:         itm.BasePriceAtOrder,
				Status:            "PENDING", // unesseory
				Notes:             itm.SpecialInstructions,
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			}

			if err := s.restaurantRepo.CreateOrderItem(txCtx, itemRecord); err != nil {
				return err
			}

			// Prepare Item for the Event
			itemEvent := restaurant_events.OrderItemEventDTO{
				MenuItemID:   itm.MenuItemID,
				MenuItemName: itm.ItemName,
				Quantity:     itm.Quantity,
				Notes:        itm.SpecialInstructions,
				Variants:     make([]restaurant_events.VariantEventDTO, 0),
				Addons:       make([]restaurant_events.AddonEventDTO, 0),
			}

			// 3. Handle Variants
			for _, vr := range itm.SelectedVariants {
				variantRecord := &restaurant_models.RestaurantOrderItemsVariants{
					ID:                    uuid.New(),
					RestaurantOrderItemID: itemRecord.ID,
					VariantID:             vr.VariantID,
					VariantGroupID:        vr.VariantGroupID,
					Name:                  vr.VariantName,
					PriceAtOrder:          vr.PriceAtOrder,
					CreatedAt:             time.Now(),
				}
				if err := s.restaurantRepo.CreateOrderItemVariant(txCtx, variantRecord); err != nil {
					return err
				}

				// Add to Event DTO
				itemEvent.Variants = append(itemEvent.Variants, restaurant_events.VariantEventDTO{
					VariantID:    vr.VariantID,
					VariantName:  vr.VariantName,
					PriceAtOrder: vr.PriceAtOrder,
				})
			}

			// 4. Handle Add-ons
			for _, adn := range itm.SelectedAddOns {
				addonRecord := &restaurant_models.RestaurantOrderItemsAddOns{
					ID:                    uuid.New(),
					RestaurantOrderItemID: itemRecord.ID,
					AddOnID:               adn.ID,
					Name:                  adn.Name,
					Quantity:              adn.Qty,
					PriceAtOrder:          adn.PriceAtOrder,
					CreatedAt:             time.Now(),
				}
				if err := s.restaurantRepo.CreateOrderItemAddon(txCtx, addonRecord); err != nil {
					return err
				}

				// Add to Event DTO
				itemEvent.Addons = append(itemEvent.Addons, restaurant_events.AddonEventDTO{
					AddOnID:      adn.ID,
					Name:         adn.Name,
					Qty:          adn.Qty,
					PriceAtOrder: adn.PriceAtOrder,
				})
			}

			// Append the fully built item DTO to the payload
			eventPayload.Items = append(eventPayload.Items, itemEvent)
		}

		// 5. Publish the clean DTO
		err := s.seroEngine.Publish(txCtx, restaurant_events.OrderCreatedInRestaurant, eventPayload)
		if err != nil {
			return err
		}

		return nil
	})
}

/* ----------------------------------- <*> ---------------------------------- */

func (s *restaurantService) GetDineSessionOrders(ctx context.Context, dineInSessionID uuid.UUID) (*restaurant_contracts.DiningSessionOrders_resp_dto, error) {

	// 1. Fetch orders from Repo
	orders, err := s.restaurantRepo.GetOrdersByDiningSessionID(ctx, dineInSessionID)
	if err != nil {
		return nil, err
	}

	resp := &restaurant_contracts.DiningSessionOrders_resp_dto{
		DiningSessionID: dineInSessionID,
		Orders:          make([]restaurant_contracts.Order_resp_dto, 0),
	}

	var runningSessionTotal float64

	for _, ord := range orders {
		orderDTO := restaurant_contracts.Order_resp_dto{
			ID:          ord.ID,
			OrderNumber: ord.OrderNumber,
			Status:      ord.Status,
			CreatedAt:   ord.CreatedAt,
			Items:       make([]restaurant_contracts.OrderItem_resp_dto, 0),
		}

		// 2. Fetch Items for this specific order
		items, _ := s.restaurantRepo.GetOrderItemsByOrderID(ctx, ord.ID)

		for _, itm := range items {
			itemDTO := restaurant_contracts.OrderItem_resp_dto{
				ID:         itm.ID,
				MenuItemID: itm.MenuItemID,
				Name:       itm.MenuItemName, // Assumes your model/join has this
				Quantity:   itm.Quantity,
				UnitPrice:  itm.UnitPrice,
				Status:     itm.Status,
				Notes:      itm.Notes,
			}

			// 3. Fetch Variants & Addons (Optimization: Use a Join in Repo instead of loops)
			variants, _ := s.restaurantRepo.GetOrderItemVariantByOrderID(ctx, itm.ID)
			for _, v := range variants {
				itemDTO.Variants = append(itemDTO.Variants, restaurant_contracts.Variant_resp_dto{
					ID:    v.VariantID,
					Name:  v.Name,
					Price: v.PriceAtOrder,
				})
			}

			addons, _ := s.restaurantRepo.GetOrderItemAddOnByOrderID(ctx, itm.ID)
			for _, a := range addons {
				itemDTO.AddOns = append(itemDTO.AddOns, restaurant_contracts.Addon_resp_dto{
					ID:       a.AddOnID,
					Name:     a.Name,
					Quantity: a.Quantity,
					Price:    a.PriceAtOrder,
				})
			}

			orderDTO.Items = append(orderDTO.Items, itemDTO)
		}

		resp.Orders = append(resp.Orders, orderDTO)
	}

	resp.SessionTotal = runningSessionTotal
	return resp, nil
}
