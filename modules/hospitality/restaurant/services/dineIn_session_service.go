package restaurantservices

import (
	"context"
	"errors"
	dbtx "sharpnr-engine/core/dbTx"
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
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

		ord := &restaurant_models.RestaurantOrders{
			ID:                        uuid.New(),
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

			item := &restaurant_models.RestaurantOrderItems{
				ID:                uuid.New(),
				RestaurantOrderID: restaurantID,
				MenuItemID:        itm.MenuItemID,
				Quantity:          itm.Quantity,
				UnitPrice:         itm.BasePriceAtOrder,
				Status:            "PENDING",
				Notes:             itm.SpecialInstructions,
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			}

			if err := s.restaurantRepo.CreateOrderItem(txCtx, item); err != nil {
				return err
			}

			if len(itm.SelectedVariants) > 0 {

				for _, vr := range itm.SelectedVariants {

					variants := &restaurant_models.RestaurantOrderItemsVariants{
						ID:                    uuid.New(),
						RestaurantOrderItemID: item.ID,
						VariantID:             vr.VariantID,
						VariantGroupID:        vr.VariantGroupID,
						Name:                  vr.VariantName,
						PriceAtOrder:          vr.PriceAtOrder,
						CreatedAt:             time.Now(),
					}

					if err := s.restaurantRepo.CreateOrderItemVariant(txCtx, variants); err != nil {
						return err
					}
				}

			}

			if len(itm.SelectedAddOns) > 0 {

				for _, adn := range itm.SelectedAddOns {
					addons := &restaurant_models.RestaurantOrderItemsAddOns{
						ID:                    uuid.New(),
						RestaurantOrderItemID: item.ID,
						AddOnID:               adn.ID,
						Name:                  adn.Name,
						Quantity:              adn.Qty,
						PriceAtOrder:          adn.PriceAtOrder,
						CreatedAt:             time.Now(),
					}

					if err := s.restaurantRepo.CreateOrderItemAddon(ctx, addons); err != nil {
						return err
					}
				}

			}
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
