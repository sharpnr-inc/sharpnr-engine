package restaurantservices

import (
	"context"
	"errors"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
	"time"

	"github.com/google/uuid"
)

/************************************************
 * This is the Service belongs to order for the *
 * All the Order belongs to Restaurant But      *
 ************************************************/

/*
This is Service Function Creating the Session of the Restaurant
Session is how many Orders and how long it took for a customer to rotate over a table.
This must be closed after the payment; till then it will remain open.
"request" is the body of the request used to create a session.
*/
func (s *restaurantService) CreateTableBookingSession(
	ctx context.Context, orgId uuid.UUID,
	request hospitalitydtos.Hospitality_Restaurant_CreateSession_REQUEST) error {

	/*****************
	 * Function Body *
	 *****************/

	// TODO : ORGID (Fetched via Middleware)--> RestaurantID --> can be fetched on the basis of the OrgID
	// *1 Need Repo function for the Getting the Table By the ID
	// *2 Then we can matched it with the RestaurantID and Verify the Status
	// Then Proceeds to Creating the Session Order and Initial Order Slice
	// and since we have the KitchenRepo Access we can find the KitchenBy ID of OrgID
	// and Create the KitchenTicket with the KitchenTicketItems

	// Which will send the Data to Kitchen Via Websocket with the SERE(Sharpnr Event Relay Engine)

	restaurant, err := s.restaurantRepo.GetRestaurantByOrgId(ctx, orgId)
	if err != nil {
		return nil
	}

	table, err := s.restaurantRepo.GetTableById(ctx, request.TableId)
	if err != nil {
		return err
	}

	if table.RestaurantID != restaurant.ID {
		return errors.New("table does not belong to this restaurant")
	}

	session := &hospitalitymodels.RestaurantTableBookingSession{
		ID:                            uuid.New(),
		RestaurantTableID:             restaurant.ID,
		RestaurantTable:               *table,
		HospitalityCustomerAndGuestID: nil,
		CustomerAndGuest:              nil,
		FromTime:                      time.Now(),
		CreatedAt:                     time.Now(),
		UpdatedAt:                     time.Now(),
	}

	if err := s.restaurantRepo.CreateRes_TableSession(ctx, session); err != nil {
		return err
	}

	// Set table as OCCUPIED
	err = s.restaurantRepo.UpdateTable(
		ctx,
		table.ID,
		map[string]any{
			"status": "OCCUPIED",
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *restaurantService) AddItemInBookingSession(
	ctx context.Context,
	orgId, sessionId uuid.UUID,
	request hospitalitydtos.Hospitality_Restaurant_AddOrderItemINTheSession_Request) error {

	// *Step 1: find the Session of Based on sessionId once we get the session
	// *Step 2: Then We will create the RestaurantTableOrder that is inside request.Order its slice or Array we will use the for loop
	// *Step 3: Then we will create KitchenTicket from the kitchenRepo.CreateKitchenTicketDBTx()
	// *Step 4: Then we create The KitchenOrderItem and adding the KitchenTickedId and kitchenTicker pointer
	// *Step 5: we will KitchenTicket's Id in the RestaurantOrderItem for Ref

	return s.restaurantRepo.WithTransaction(ctx, func(txRepo restaurantrepo.RestaurantRepository) error {
		// ------------------------------------------------------------
		// STEP 1: Get Booking Session
		// ------------------------------------------------------------
		session, err := txRepo.GetBookingSessionById(ctx, sessionId)
		if err != nil {
			return err
		}

		kitchen, err := s.kitchenRepo.GetKitchenByOrgId(ctx, orgId)
		if err != nil {
			return nil
		}

		// ------------------------------------------------------------
		// STEP 2: Create RestaurantOrderItems
		// ------------------------------------------------------------

		createdOrderItems := make([]hospitalitymodels.RestaurantTableOrder, 0)

		for _, item := range request.Order {

			order := hospitalitymodels.RestaurantTableOrder{
				ID:                              uuid.New(),
				MenuItemID:                      item.MenuItemId,
				VariantID:                       item.VariantId,
				AddonID:                         item.AddonOnId,
				Quantity:                        item.Quantity,
				Notes:                           item.Notes,
				RestaurantTableBookingSessionID: session.ID,
				RestaurantTableID:               uuid.MustParse(request.TableId),
			}

			saved, err := txRepo.CreateRestaurantTableOrder(ctx, &order)
			if err != nil {
				return nil
			}

			createdOrderItems = append(createdOrderItems, *saved)
		}

		// ------------------------------------------------------------
		// STEP 3: Create Kitchen Ticket
		// ------------------------------------------------------------

		ticketNo, err := s.kitchenRepo.GetNextTicketNumber(ctx, kitchen.ID)
		if err != nil {
			return err
		}

		kitchenTicket := hospitalitymodels.KitchenTicket{
			ID:             uuid.New(),
			OrganizationID: orgId,
			KitchenID:      kitchen.ID,
			Kitchen:        *kitchen,
			OrderID:        sessionId,
			Status:         "PENDING",
			TicketNo:       ticketNo,
		}

		savedTicket, err := s.kitchenRepo.CreateKitchenTicketDBTx(ctx, &kitchenTicket)
		if err != nil {
			return err
		}

		// ------------------------------------------------------------
		// STEP 4: Create Kitchen Ticket Items
		// ------------------------------------------------------------

		for i, orderItem := range createdOrderItems {

			ktItem := hospitalitymodels.KitchenTicketItem{
				ID:              uuid.New(),
				KitchenTicketID: savedTicket.ID,
				OrderItemID:     orderItem.ID,
				MenuItemID:      orderItem.MenuItemID,
				Ticket:          kitchenTicket,
				OrderType:       "RESTAURANT",
				Quantity:        orderItem.Quantity,
				Status:          "PENDING",
			}

			savedKTItem, err := s.kitchenRepo.CreateKitchenTicketItemDBTx(ctx, &ktItem)
			if err != nil {
				return nil
			}

			err = txRepo.UpdateRestaurantOrderItemKotReference(
				ctx,
				orderItem.ID,
				savedTicket.ID,
				savedKTItem.ID,
			)
			if err != nil {
				return err
			}

			// attach back to local array if you need
			createdOrderItems[i].KitchenTicketID = &savedTicket.ID
			createdOrderItems[i].KitchenTicketItemID = &savedKTItem.ID
		}
		return nil

	})

}

func (s *restaurantService) GetResTableBookingSessionByTableId(
	ctx context.Context,
	orgId, tableId uuid.UUID,
) (hospitalitydtos.Hospitality_Restaurant_GetBookingSession_Response, error) {

	session, err := s.restaurantRepo.GetBookingSessionByTableId(ctx, tableId)
	if err != nil {
		return hospitalitydtos.Hospitality_Restaurant_GetBookingSession_Response{}, err
	}

	// session may be nil → no active session yet
	return hospitalitydtos.Hospitality_Restaurant_GetBookingSession_Response{
		Session: session,
	}, nil
}
