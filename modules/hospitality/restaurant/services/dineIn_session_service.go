package restaurantservices

import (
	"context"
	"errors"
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
			RestaurantID:              restaurantID,
			RestaurantDiningSessionID: &dineInSessionID,
			OrderNumber:               "1",
			Status:                    "PENDING",
			CreatedAt:                 time.Now(),
			UpdatedAt:                 time.Now(),
		}

		if err := s.restaurantRepo.CreateOrder(txCtx, ord); err != nil {
			return err
		}

		for _, itm := range req.Items {
			itemID := uuid.New()

			// Build the Item Record for the DB
			itemRecord := &restaurant_models.RestaurantOrderItems{
				ID:                itemID,
				RestaurantOrderID: ord.ID,
				MenuItemID:        itm.MenuItemID,
				Quantity:          itm.Quantity,
				UnitPrice:         itm.BasePriceAtOrder,
				Status:            "PENDING",
				Notes:             itm.SpecialInstructions,
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			}

			if err := s.restaurantRepo.CreateOrderItem(txCtx, itemRecord); err != nil {
				return err
			}

			// Prepare Item for the Event
			itemEvent := restaurant_events.OrderItemEventDTO{
				ID:         itemID,
				MenuItemID: itm.MenuItemID,
				Quantity:   itm.Quantity,
				Notes:      itm.SpecialInstructions,
				Variants:   make([]restaurant_events.VariantEventDTO, 0),
				Addons:     make([]restaurant_events.AddonEventDTO, 0),
			}

			// 3. Handle Variants
			for _, vr := range itm.SelectedVariants {
				variantRecord := &restaurant_models.RestaurantOrderItemsVariants{
					ID:                    uuid.New(),
					RestaurantOrderItemID: itemID,
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
					RestaurantOrderItemID: itemID,
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

func (s *restaurantService) GetDineSessionOrders(ctx context.Context, dineInSessionID uuid.UUID) (interface{}, error) {
	// 1. Fetch all orders linked to this session
	orders, err := s.restaurantRepo.GetOrdersByDiningSessionID(ctx, dineInSessionID)

	// If no orders exist for this session, return nil/null as requested
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var ordersWithItems []OrderWithItemsResponse

	// 2. Loop through each order to find its specific items
	for _, order := range orders {
		items, err := s.restaurantRepo.GetOrderItemsByOrderID(ctx, order.ID)

		// Handle "No Items" case gracefully
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// If an order exists but has no items, we treat items as empty
				items = []*restaurant_models.RestaurantOrderItems{}
			} else {
				return nil, err
			}
		}

		ordersWithItems = append(ordersWithItems, OrderWithItemsResponse{
			Order: order,
			Items: items,
		})
	}

	// 3. Return the grouped response
	return DineSessionOrdersResponse{
		SessionID: dineInSessionID,
		Orders:    ordersWithItems,
	}, nil
}

type OrderWithItemsResponse struct {
	Order *restaurant_models.RestaurantOrders       `json:"order"`
	Items []*restaurant_models.RestaurantOrderItems `json:"items"`
}

type DineSessionOrdersResponse struct {
	SessionID uuid.UUID                `json:"session_id"`
	Orders    []OrderWithItemsResponse `json:"orders"`
}
